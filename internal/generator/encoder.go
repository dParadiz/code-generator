package generator

import (	
	"plugin"
)

type Encoder interface {
	Encode(input  string) (interface{}, error)
}

func loadEncoder(name string) Encoder, error {
	plugin, err := plugin.Open(name)

	if err != nil {
		return err
	}

	symCheck, err := plugin.Lookup("Encode")

	if err != nil {
		return err
	}

	var enc = Encoder
	encoder, ok := symCheck.(enc)
	if !ok {
		return errors.New("unexpected type from module symbol")
	}

	return encoder

}
