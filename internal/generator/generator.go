package generator

import "github.com/dparadiz/code-generator/internal/renderer"

type Context struct {
	EncoderName   string
	EncoderConfig string
	DecoderName   string
	DecoderConfig string
	OutputFolder  string
}

func (c *Context) Generate() {
	encoderImplementation, err := loadEncoder(c.EncoderName)
	if err != nil {
		panic(err)
	}

	decoderImplementation, err := loadDecoder(c.DecoderName)
	if err != nil {
		panic(err)
	}

	data, err := encoderImplementation.Encode(c.EncoderConfig)
	if err != nil {
		panic(err)
	}

	decoderContext := new(DecoderContext)
	decoderContext.ConfigSource = c.DecoderConfig
	decoderContext.Data = data

	stack := new(renderer.Stack)

	err = decoderImplementation.Decode(decoderContext, stack)
	if err != nil {
		panic(err)
	}

	renderer.Process(c.OutputFolder, stack)
}
