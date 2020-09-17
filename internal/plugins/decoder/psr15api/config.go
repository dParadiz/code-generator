package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

var cfg config

type config struct {
	ModelTemplate              string `json:"model_template"`
	RequestHandlerTemplate     string `json:"request_handler_template"`
	Namespace                  string `json:"namespace"`
	ModelNamespace             string `json:"model_namespace"`
	PathHandlerNamespace       string `json:"path_handler_namespace" default:"PathHandler"`
	OperationInterfaceTemplate string `json:"operation_interface_template"`
}

func (c *config) getModelNamespace() string {
	var namespaceBuilder strings.Builder

	if c.Namespace != "" {
		namespaceBuilder.WriteString(c.Namespace + "\\")
	}

	if c.ModelNamespace == "" {
		namespaceBuilder.WriteString("Model")
	} else {
		namespaceBuilder.WriteString(c.ModelNamespace)
	}

	return namespaceBuilder.String()
}

func (c *config) getPathHandlerNamespace() string {

	var namespaceBuilder strings.Builder

	if c.Namespace != "" {
		namespaceBuilder.WriteString(c.Namespace + "\\")
	}

	if c.PathHandlerNamespace == "" {
		namespaceBuilder.WriteString("Path")
	} else {
		namespaceBuilder.WriteString(c.PathHandlerNamespace)
	}

	return namespaceBuilder.String()
}

func loadConfig(configPath string) error {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cfg)
	return nil
}
