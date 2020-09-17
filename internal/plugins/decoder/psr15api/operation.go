package main

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type Parameter struct {
	Name          string
	Required      bool
	DocType       string
	Type          string
	Default       interface{}
	FullClassName string
	Last          bool
}

type Operation struct {
	Name       string
	Namespace  string
	Method     string
	Parameters []Parameter
	Imports    []string
	Last       bool
}

func (o *Operation) setName(operationId string) {
	name := strings.Title(operationId)
	if name == "" {
		name = strings.Title(o.Method)
	}

	o.Name = name + "Interface"
}

// TODO add common parameter
// TODO parameter style and explode support
// TODO parameter context support
func (o *Operation) setParameters(operation *openapi3.Operation) {

	for _, parameter := range operation.Parameters {
		param := Parameter{
			Name:     parameter.Value.Name,
			Type:     getType(parameter.Value.Schema),
			DocType:  getDocType(parameter.Value.Schema),
			Required: parameter.Value.Required,
			Default:  parameter.Value.Schema.Value.Default,
			Last:     false,
		}

		if !param.Required && param.Default == nil {
			param.Default = "null"
			param.Type = "?" + param.Type
			param.DocType += "|null"
		}

		o.Parameters = append(o.Parameters, param)

	}

	if operation.RequestBody != nil && operation.RequestBody.Value.Content["application/json"].Schema != nil {
		// TODO support for other media types
		schema := operation.RequestBody.Value.Content["application/json"].Schema

		paramName := getType(schema)
		// making first character lowercase
		var b strings.Builder
		b.WriteString(strings.ToLower(string(paramName[0])))
		b.WriteString(paramName[1:])

		param := Parameter{
			Required: operation.RequestBody.Value.Required,
			Name:     b.String(),
			Type:     getType(schema),
			DocType:  getDocType(schema),
			Default:  schema.Value.Default,
			Last:     false,
		}

		if !param.Required && param.Default == nil {
			param.Default = "null"
			param.Type = "?" + param.Type
			param.DocType += "|null"
		}

		o.Imports = append(o.Imports, cfg.getModelNamespace()+"\\"+param.Type)

		o.Parameters = append(o.Parameters, param)
	}

	if len(o.Parameters) > 0 {
		// bring required parameters in front
		sort.SliceStable(o.Parameters, func(i, j int) bool {
			return o.Parameters[i].Required
		})

		o.Parameters[len(o.Parameters)-1].Last = true
	}
}

func (o Operation) getClassNamespace() string {
	return o.Namespace + "\\" + o.Name
}
