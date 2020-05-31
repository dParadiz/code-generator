package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/dparadiz/code-generator/internal/generator"
	"github.com/dparadiz/code-generator/internal/renderer"
)

type config struct {
	Template string `json:"template"`
}

type decoderContext string

func (dc decoderContext) Decode(c *generator.DecoderContext, stack *renderer.Stack) error {

	cfg, err := loadConfig(c.ConfigSource)
	if err != nil {
		return err
	}

	stackItem := new(renderer.StackItem)
	stackItem.Output = "GeneratedFile.md"
	stackItem.Template = cfg.Template
	stackItem.TemplateData = c.Data

	stack.Push(stackItem)

	return nil
}

func loadConfig(configPath string) (config, error) {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return config{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config config
	json.Unmarshal(byteValue, &config)

	return config, nil
}

// export as symbol Decoder
var Decoder decoderContext
