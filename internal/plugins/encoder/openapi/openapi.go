package main

import "github.com/getkin/kin-openapi/openapi3"

type encoderContext string

func (ec encoderContext) Encode(input string) (interface{}, error) {

	return openapi3.NewSwaggerLoader().LoadSwaggerFromFile(input)
}

// Encoder exported symbol
var Encoder encoderContext
