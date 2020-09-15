package main

import "strings"

type PropertyValidator struct {
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
	Value    string
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

type Model struct {
	Name               string
	Namespace          string
	OptionalProperties []Property
	RequiredProperties []Property
	Constant           []Property
}
