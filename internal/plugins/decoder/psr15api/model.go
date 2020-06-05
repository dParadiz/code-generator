package main

type Property struct {
	Name string
	Type string
	DocType string
	Value interface{}
	Visibility string 
}

type Model struct {
	Name string
	Namespace string
	OptionalProperties []Property
	RequiredProperties []Property
	Constant		[]Property
}