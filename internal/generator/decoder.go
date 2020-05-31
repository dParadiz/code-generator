package generator

import (
	"plugin"

	"github.com/dparadiz/code-generator/internal/generator/"
)

type DecoderContext struct {
	ConfigSource string,
	Data interface{}
}

type Decoder interface {
	Decode(context *DecoderContext, stack *renderer.Stack) error
}

func loadDecorator(name string) Decoder, error {
	plugin, err := plugin.Open(name)

	if err != nil {
		return err
	}

	symCheck, err := plugin.Lookup("Decode")

	if err != nil {
		return err
	}

	var enc = Decoder
	encoder, ok := symCheck.(enc)
	if !ok {
		return errors.New("unexpected type from module symbol")
	}

	return encoder

}