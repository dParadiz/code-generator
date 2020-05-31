package generator

type Context struct {
	EncoderName string
	EncoderConfig string
	DecoderName string
	DecoderConfig string
	OutputFolder string
}

func (context *Context) generate()  {
	encoderImplementation, err := encoder.loadEncoder(context.EncoderName)
	if err !== nil {
		panic(err)
	}
	decoderImplementation, err := decoder.loadDecoder(context.DecoderName)
	if err !== nil {
		panic(err)
	}
	
	data, err := encoderImplementation.Encode(context.EncoderConfig);
	if err != nil {
		panic(err)
	}
	
	decoderContext = DecoderContext{ConfigSource : concontext.DecoderConfig, Data: data}
	stack := new(renderer.Stack);

	err := decoderImplementation.Decode(decoderContext, stack);
	if err != nil {
		panic(err)
	}	

	renderer.Process(context.outputFolder, stack);	
}