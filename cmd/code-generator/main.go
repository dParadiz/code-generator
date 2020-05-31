package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dparadiz/code-generator/internal/generator"
)

func main() {

	outputFolder := flag.String("output", ".", "Output folder")
	encoderConfig := flag.String("encoder-config", "", "Encoder specific config")

	encoderName := flag.String("encoder", "", "Encoder to be used")

	decoderName := flag.String("decoder", "", "Decoder to be used")
	decoderConfig := flag.String("decoder-config", "", "Decoder specific config")
	flag.Parse()

	if *encoderName == "" {
		fmt.Println("Encoder name can not be empty")
		os.Exit(1)

	}

	context := generator.Context{
		EncoderName:   *encoderName,
		EncoderConfig: *encoderConfig,
		DecoderName:   *decoderName,
		DecoderConfig: *decoderConfig,
		OutputFolder:  *outputFolder,
	}

	context.Generate()
}
