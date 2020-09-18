package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/dparadiz/code-generator/internal/generator"
	"github.com/dparadiz/code-generator/internal/renderer"
)

type decoderContext string

type phpClass interface {
	getClassNamespace() string
}

func (dc decoderContext) Decode(c *generator.DecoderContext, stack *renderer.Stack) error {

	err := loadConfig(c.ConfigSource)
	if err != nil {
		return err
	}

	var openapi *openapi3.Swagger
	openapi, ok := c.Data.(*openapi3.Swagger)
	if !ok {
		return errors.New("Can work only with openapi 3.0 data")
	}

	for schemaName, schema := range openapi.Components.Schemas {

		if schema.Value.Type != "object" {
			continue
		}

		model := Model{
			Name:      schemaName,
			Namespace: cfg.getModelNamespace(),
		}

		model.populateFrom(schema)

		stackItem := new(renderer.StackItem)
		stackItem.Output = getPsr4AutoloadFilename(cfg.Namespace, model)
		stackItem.Template = cfg.ModelTemplate
		stackItem.TemplateData = model

		stack.Push(stackItem)

	}

	for path, pathItem := range openapi.Paths {

		camelcasedPath := toCamelCase(path)
		requestHandler := new(RequestHandler)
		requestHandler.Name = "RequestHandler"
		requestHandler.Namespace = cfg.getPathHandlerNamespace() + "\\" + camelcasedPath
		requestHandler.setOperations(pathItem)

		stackItem := new(renderer.StackItem)
		stackItem.Output = getPsr4AutoloadFilename(cfg.Namespace, requestHandler)
		stackItem.Template = cfg.RequestHandlerTemplate
		stackItem.TemplateData = requestHandler

		stack.Push(stackItem)

		for _, operation := range requestHandler.Operations {
			stackItem := new(renderer.StackItem)
			stackItem.Output = getPsr4AutoloadFilename(cfg.Namespace, operation)
			stackItem.Template = cfg.OperationInterfaceTemplate
			stackItem.TemplateData = operation

			stack.Push(stackItem)
		}

	}

	return nil
}

func getType(schemaType *openapi3.SchemaRef) string {
	if schemaType.Value.AllOf != nil {
		return schemaType.Ref[strings.LastIndex(schemaType.Ref, "/")+1:]
	}

	switch propertyType := schemaType.Value.Type; propertyType {
	case "number":
		return "float"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	case "object":
		return schemaType.Ref[strings.LastIndex(schemaType.Ref, "/")+1:]
	case "array":
		return "array"
	case "string":
		return "string"
	default:
		panic(fmt.Sprintf("Unsupported type " + propertyType))
	}
}

func getDocType(schema *openapi3.SchemaRef) string {
	pType := getType(schema)

	if strings.Compare(pType, "array") == 0 {
		if schema.Value.Items.Ref != "" {
			return "[]" + schema.Value.Items.Ref[strings.LastIndex(schema.Value.Items.Ref, "/")+1:]
		}
		return "[]" + schema.Value.Items.Value.Type
	}

	return pType
}

func getPsr4AutoloadFilename(namespace string, model phpClass) string {
	classNamespace := model.getClassNamespace()

	var stringBuilder strings.Builder

	if namespace == "" {
		stringBuilder.WriteString("/")
	} else {
		classNamespace = strings.ReplaceAll(classNamespace, namespace, "")
	}

	stringBuilder.WriteString(strings.ReplaceAll(classNamespace, "\\", "/"))
	stringBuilder.WriteString(".php")

	return stringBuilder.String()
}

func toCamelCase(path string) string {

	path = string(regexp.MustCompile(`[^A-Za-z\d]`).ReplaceAll([]byte(path), []byte(" ")))
	path = strings.Title(path)
	return string(regexp.MustCompile(` `).ReplaceAll([]byte(path), []byte("")))
}

// Decoder symbol export
var Decoder decoderContext
