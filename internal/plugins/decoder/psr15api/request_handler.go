package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type RequestHandler struct {
	Name       string `default:"RequestHandler"`
	Operations []Operation
	Imports    []string
	Namespace  string
}

func (r *RequestHandler) addOperation(operation Operation) {
	r.Operations = append(r.Operations, operation)
	r.Imports = append(r.Imports, operation.Namespace+"\\"+operation.Name)
}

func (r *RequestHandler) setOperations(path *openapi3.PathItem) {
	if path.Get != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "get",
			Last:      false,
		}
		operation.setName(path.Get.OperationID)

		operation.setParameters(path.Get)
		operation.setResponses(path.Get)

		r.addOperation(operation)

	}

	if path.Delete != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "delete",
			Last:      false,
		}

		operation.setName(path.Delete.OperationID)

		operation.setParameters(path.Delete)
		operation.setResponses(path.Delete)

		r.addOperation(operation)
	}

	if path.Post != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "post",
			Last:      false,
		}

		operation.setName(path.Post.OperationID)

		operation.setParameters(path.Post)
		operation.setResponses(path.Post)

		r.addOperation(operation)
	}

	if path.Put != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "put",
			Last:      false,
		}

		operation.setName(path.Put.OperationID)

		operation.setParameters(path.Put)
		operation.setResponses(path.Put)

		r.addOperation(operation)
	}

	if path.Patch != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "patch",
			Last:      false,
		}

		operation.setName(path.Patch.OperationID)

		operation.setParameters(path.Patch)
		operation.setResponses(path.Patch)

		r.addOperation(operation)
	}

	if path.Options != nil {
		operation := Operation{
			Namespace: r.getOperationNamespace(),
			Method:    "options",
			Last:      false,
		}

		operation.setName(path.Options.OperationID)

		operation.setParameters(path.Options)
		operation.setResponses(path.Options)

		r.addOperation(operation)
	}

	if len(r.Operations) > 0 {
		r.Operations[len(r.Operations)-1].Last = true
	}

}
func (r RequestHandler) getOperationNamespace() string {
	return r.Namespace + "\\Operation"
}

func (r RequestHandler) getClassNamespace() string {
	return r.Namespace + "\\" + r.Name
}
