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

		// prepare required map
		required := make(map[string]bool)
		for _, name := range schema.Value.Required {
			required[name] = true

		}

		for propertyName, property := range schema.Value.Properties {

			var modelProperty = new(Property)
			modelProperty.PropertyName = propertyName
			modelProperty.Value = property.Value.Default
			modelProperty.Visibility = "private"
			modelProperty.Nullable = property.Value.Nullable

			modelProperty.Type = getType(property)
			modelProperty.DocType = getType(property)
			modelProperty.Last = false

			setPropertyValidator(modelProperty, property)

			if required[propertyName] {
				model.RequiredProperties = append(model.RequiredProperties, *modelProperty)
			} else {
				model.OptionalProperties = append(model.OptionalProperties, *modelProperty)
			}
		}

		if len(model.RequiredProperties) > 0 {
			model.RequiredProperties[len(model.RequiredProperties)-1].Last = true
		}

		if len(model.OptionalProperties) > 0 {
			model.OptionalProperties[len(model.OptionalProperties)-1].Last = true
		}

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

func setPropertyValidator(property *Property, schema *openapi3.SchemaRef) {
	if schema.Value.Min != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.Min))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.Min)
		}

		property.Validators = append(property.Validators, PropertyValidator{Min: true, Value: value})
	}

	if schema.Value.Max != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.Max))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.Max)
		}
		property.Validators = append(property.Validators, PropertyValidator{Max: true, Value: value})
	}

	if schema.Value.MultipleOf != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.MultipleOf))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.MultipleOf)
		}
		property.Validators = append(property.Validators, PropertyValidator{MultipleOf: true, Value: value})
	}

	if schema.Value.MaxLength != nil {
		property.Validators = append(property.Validators, PropertyValidator{MaxLength: true, Value: fmt.Sprintf("%d", *schema.Value.MaxLength)})
	}

	if schema.Value.MinLength > 0 {
		property.Validators = append(property.Validators, PropertyValidator{MinLength: true, Value: fmt.Sprintf("%d", schema.Value.MinLength)})
	}

	if schema.Value.MinItems > 0 {
		property.Validators = append(property.Validators, PropertyValidator{MinItems: true, Value: fmt.Sprintf("%d", schema.Value.MinItems)})
	}

	if schema.Value.MaxItems != nil {
		property.Validators = append(property.Validators, PropertyValidator{MaxItems: true, Value: fmt.Sprintf("%d", *schema.Value.MaxItems)})
	}
}

// Decoder symbol export
var Decoder decoderContext
