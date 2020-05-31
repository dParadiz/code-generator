package main

import (
	"flag"

	"github.com/dparadiz/code-generator/internal/generator"
)

func main() {

	outputFolder := flag.String("output", "", "Output folder")
	encoderConfig := flag.String("encoder-config", "", "Encoder specific config")
	encoderName := flag.String("encoder", "", "Encoder to be used")
	decoderName := flag.String("decoder", "", "Decoder to be used")
	decoderConfig := flag.String("decoder-config", "", "Decoder specific config")

	context := generator.Context{
		EncoderName:   *encoderName,
		EncoderConfig: *encoderConfig,
		DecoderName:   *decoderName,
		DecoderConfig: *decoderConfig,
		OutputFolder:  *outputFolder,
	}

	context.Generate()
}
