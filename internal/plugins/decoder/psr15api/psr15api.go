package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/dparadiz/code-generator/internal/generator"
	"github.com/dparadiz/code-generator/internal/renderer"
)

type config struct {
	TemplateFolder string `json:"templateFolder"`
}

type decoderContext string

func (dc decoderContext) Decode(c *generator.DecoderContext, stack *renderer.Stack) error {

	cfg, err := loadConfig(c.ConfigSource)
	if err != nil {
		return err
	}

	fmt.Printf("Configuration data: %v\n", cfg)

	var openapi *openapi3.Swagger
	openapi, ok := c.Data.(*openapi3.Swagger)
	if !ok {
		return errors.New("Can work only with openapi 3.0 data")
	}

	for schemaName, schema := range openapi.Components.Schemas {
		fmt.Printf("%s\n", schemaName)
		fmt.Printf("\treqired %v\n", schema.Value.Required)
		fmt.Println("\tProperties:")
		for propertyName, property := range schema.Value.Properties {
			fmt.Printf("\t\t%s %s\n", propertyName, property.Value.Type)
		}

	}

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

// Decoder symbol export
var Decoder decoderContext
