package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type Model struct {
	Name               string
	Namespace          string
	OptionalProperties []Property
	RequiredProperties []Property
	Constant           []Property
}

func (m *Model) populateFrom(schema *openapi3.SchemaRef) {

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

		if property.Value.Nullable {
			modelProperty.Type = "?" + getType(property)
		} else {
			modelProperty.Type = getType(property)
		}

		modelProperty.DocType = getDocType(property)
		modelProperty.Last = false
		modelProperty.setValidators(property)

		if required[propertyName] {
			m.RequiredProperties = append(m.RequiredProperties, *modelProperty)
		} else {
			m.OptionalProperties = append(m.OptionalProperties, *modelProperty)
		}
	}

	if len(m.RequiredProperties) > 0 {
		m.RequiredProperties[len(m.RequiredProperties)-1].Last = true
	}

	if len(m.OptionalProperties) > 0 {
		m.OptionalProperties[len(m.OptionalProperties)-1].Last = true
	}
}

func (m Model) getClassNamespace() string {
	return m.Namespace + "\\" + m.Name
}
