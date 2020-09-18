package main

import (
	"regexp"
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

type Return struct {
	Type    string
	DocType string
}

type Exception struct {
	ContentType      string
	ContentName      string
	Imports          []string
	Name             string
	Namespace        string
	ExceptionType    string
	ExceptionMessage string
	ExceptionCode    string
	HasContent       bool
}

func (e Exception) getClassNamespace() string {
	return e.Namespace + "\\" + e.Name
}

type Operation struct {
	Name          string
	Namespace     string
	Method        string
	Parameters    []Parameter
	ReturnType    string
	ReturnDocType string
	ReturnCode    string
	Exceptions    []Exception
	Imports       []string
	Last          bool
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
func (o *Operation) setResponses(operation *openapi3.Operation) {

	for code, response := range operation.Responses {
		matched200, _ := regexp.MatchString(`2.*`, code)
		matched300, _ := regexp.MatchString(`3.*`, code)
		matched400, _ := regexp.MatchString(`4.*`, code)
		matched500, _ := regexp.MatchString(`5.*`, code)

		var cType, cDocType string

		if response.Value.Content["application/json"] != nil {
			cType = getType(response.Value.Content["application/json"].Schema)
			cDocType = getDocType(response.Value.Content["application/json"].Schema)
		}

		if matched200 {
			o.ReturnCode = code
			if cType == "" {
				o.ReturnType = "void"
			} else {
				o.ReturnType = cType

				if cType == "array" {
					cType = strings.ReplaceAll(cDocType, "[]", "")
				}

				skipTypes := map[string]bool{
					"int":    true,
					"bool":   true,
					"float":  true,
					"string": true,
					"array":  true,
				}

				if !skipTypes[cType] {
					o.Imports = append(o.Imports, cfg.getModelNamespace()+"\\"+cType)
				}

			}

			if cDocType == "" {
				o.ReturnDocType = "void"
			} else {
				o.ReturnDocType = cDocType
			}

		} else {

			exception := Exception{
				ExceptionMessage: *response.Value.Description,
				ExceptionCode:    "200",
				HasContent:       false,
				Name:             "Exception",
				Namespace:        o.Namespace + "\\Exception",
			}

			if matched300 {
				panic("Response code 300 are not supported")
			}
			if matched400 {
				exception.ExceptionCode = code
				exception.Name = "ClientError" + code + "Exception"
			}

			if matched500 {
				exception.ExceptionCode = code
				exception.Name = "ServerError" + code + "Exception"
			}

			if cType != "" {
				var b strings.Builder
				b.WriteString(strings.ToLower(string(cType[0])))
				b.WriteString(cType[1:])

				exception.ContentName = b.String()
				exception.ContentType = cType
				exception.HasContent = true
			}

			o.Exceptions = append(o.Exceptions, exception)
			o.Imports = append(o.Imports, exception.getClassNamespace())

		}

	}
}
func (o Operation) getClassNamespace() string {
	return o.Namespace + "\\" + o.Name
}
