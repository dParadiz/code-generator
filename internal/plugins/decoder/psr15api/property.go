package main

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type PropertyValidator struct {
	Value string
	// Numbers
	Min        bool
	Max        bool
	MultipleOf bool
	// String
	MinLength bool
	MaxLength bool
	// Arrays
	MinItems bool
	MaxItems bool
	// Object
	MinProps bool
	MaxProps bool
}

type Property struct {
	PropertyName string
	Type         string
	DocType      string
	Value        interface{}
	Visibility   string
	Nullable     bool
	Last         bool
	Validators   []PropertyValidator
}

func (p Property) CName() string {
	return strings.Title(p.PropertyName)
}

func (p *Property) setValidators(schema *openapi3.SchemaRef) {
	if schema.Value.Min != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.Min))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.Min)
		}

		p.Validators = append(p.Validators, PropertyValidator{Min: true, Value: value})
	}

	if schema.Value.Max != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.Max))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.Max)
		}
		p.Validators = append(p.Validators, PropertyValidator{Max: true, Value: value})
	}

	if schema.Value.MultipleOf != nil {
		var value string
		if schema.Value.Type == "integer" {
			value = fmt.Sprintf("%d", int64(*schema.Value.MultipleOf))
		} else {
			value = fmt.Sprintf("%f", *schema.Value.MultipleOf)
		}
		p.Validators = append(p.Validators, PropertyValidator{MultipleOf: true, Value: value})
	}

	if schema.Value.MaxLength != nil {
		p.Validators = append(p.Validators, PropertyValidator{MaxLength: true, Value: fmt.Sprintf("%d", *schema.Value.MaxLength)})
	}

	if schema.Value.MinLength > 0 {
		p.Validators = append(p.Validators, PropertyValidator{MinLength: true, Value: fmt.Sprintf("%d", schema.Value.MinLength)})
	}

	if schema.Value.MinItems > 0 {
		p.Validators = append(p.Validators, PropertyValidator{MinItems: true, Value: fmt.Sprintf("%d", schema.Value.MinItems)})
	}

	if schema.Value.MaxItems != nil {
		p.Validators = append(p.Validators, PropertyValidator{MaxItems: true, Value: fmt.Sprintf("%d", *schema.Value.MaxItems)})
	}
}
