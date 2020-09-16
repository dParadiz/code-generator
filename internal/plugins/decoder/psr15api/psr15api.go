package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/dparadiz/code-generator/internal/generator"
	"github.com/dparadiz/code-generator/internal/renderer"
)

type decoderContext string

func (dc decoderContext) Decode(c *generator.DecoderContext, stack *renderer.Stack) error {

	cfg, err := loadConfig(c.ConfigSource)
	if err != nil {
		return err
	}

	var openapi *openapi3.Swagger
	openapi, ok := c.Data.(*openapi3.Swagger)
	if !ok {
		return errors.New("Can work only with openapi 3.0 data")
	}

	modelNamespace := modelNamespace(cfg.Namespace)

	for schemaName, schema := range openapi.Components.Schemas {

		if schema.Value.Type != "object" {
			continue
		}

		model := new(Model)
		model.Name = schemaName
		model.Namespace = modelNamespace
		model.populateFrom(schema)

		stackItem := new(renderer.StackItem)
		stackItem.Output = modelOutputFile(schemaName)
		stackItem.Template = cfg.ModelTemplate
		stackItem.TemplateData = model

		stack.Push(stackItem)

	}

	return nil
}

func getType(schemaType *openapi3.SchemaRef) string {
	switch propertyType := schemaType.Value.Type; propertyType {
	case "number":
		if schemaType.Value.Nullable {
			return "?float"
		}
		return "float"
	case "integer":
		if schemaType.Value.Nullable {
			return "?int"
		}
		return "int"
	case "boolean":
		return "bool"
	default:
		if schemaType.Value.Nullable {
			return fmt.Sprintf("?%s", propertyType)
		}
		return propertyType
	}
}

func modelNamespace(baseNamespace string) string {
	var modelNamespace strings.Builder
	modelNamespace.WriteString(baseNamespace)
	modelNamespace.WriteString("\\Model")

	return modelNamespace.String()
}

func modelOutputFile(schemaName string) string {
	var outputFileName strings.Builder
	outputFileName.WriteString("/Model/")
	outputFileName.WriteString(schemaName)
	outputFileName.WriteString(".php")
	return outputFileName.String()
}

// Decoder symbol export
var Decoder decoderContext
