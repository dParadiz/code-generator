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

	symCheck, err := plugin.Lookup("Decoder")

	if err != nil {
		return nil, err
	}

	var decoder Decoder
	decoder, ok := symCheck.(Decoder)
	if !ok {
		return nil, errors.New("Unexpected  type from decoder symbol")
	}

	return decoder, nil

}
