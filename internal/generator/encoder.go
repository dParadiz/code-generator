package generator

import (
	"errors"
	"plugin"
)

type Encoder interface {
	Encode(input string) (interface{}, error)
}

func loadEncoder(name string) (Encoder, error) {
	plugin, err := plugin.Open(name)

	if err != nil {
		return nil, err
	}

	symEncoder, err := plugin.Lookup("Encoder")

	if err != nil {
		return nil, err
	}

	var encoder Encoder
	encoder, ok := symEncoder.(Encoder)
	if !ok {
		return nil, errors.New("Unexpected type from encoder symbol")
	}

	return encoder, nil

}
