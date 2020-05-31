package generator

import (
	"errors"
	"plugin"

	"github.com/dparadiz/code-generator/internal/renderer"
)

type DecoderContext struct {
	ConfigSource string
	Data         interface{}
}

type Decoder interface {
	Decode(context *DecoderContext, stack *renderer.Stack) error
}

func loadDecoder(name string) (Decoder, error) {
	plugin, err := plugin.Open(name)

	if err != nil {
		return nil, err
	}

	symCheck, err := plugin.Lookup("Decode")

	if err != nil {
		return nil, err
	}

	var encoder Decoder
	encoder, ok := symCheck.(Decoder)
	if !ok {
		return nil, errors.New("unexpected type from module symbol")
	}

	return encoder, nil

}
