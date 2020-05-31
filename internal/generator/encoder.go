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

	symCheck, err := plugin.Lookup("Encode")

	if err != nil {
		return nil, err
	}

	var encoder Encoder
	encoder, ok := symCheck.(Encoder)
	if !ok {
		return nil, errors.New("unexpected type from module symbol")
	}

	return encoder, nil

}
