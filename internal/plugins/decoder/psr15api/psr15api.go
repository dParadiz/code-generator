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

	for schemaName, schema := range openapi.Components.Schemas {

		if schema.Value.Type != "object" {
			continue
		}

		var outputFileName strings.Builder
		outputFileName.WriteString("/Model/")
		outputFileName.WriteString(schemaName)
		outputFileName.WriteString(".php")

		model := new(Model)
		model.Name = schemaName
		model.Namespace = "Model"

		//fmt.Printf("\treqired %v\n", schema.Value.Required)

		for propertyName, property := range schema.Value.Properties {

			var modelProperty = new(Property)
			modelProperty.Name = propertyName
			modelProperty.Value = property.Value.Default
			modelProperty.Visibility = "private"
			modelProperty.Type = getType(property)
			modelProperty.DocType = getType(property)

			model.OptionalProperties = append(model.OptionalProperties, *modelProperty)
		}
		fmt.Println(outputFileName.String())
		stackItem := new(renderer.StackItem)
		stackItem.Output = outputFileName.String()
		stackItem.Template = cfg.ModelTemplate
		stackItem.TemplateData = model

		stack.Push(stackItem)

	}

	return nil
}

func getType(schemaType *openapi3.SchemaRef) string {
	switch propertyType := schemaType.Value.Type; propertyType {
	case "number":
		return "float"
	case "integer":
		return "int"
	case "boolean":
		return "bool"
	default:
		return propertyType
	}
}

// Decoder symbol export
var Decoder decoderContext
