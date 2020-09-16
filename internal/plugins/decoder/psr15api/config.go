package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	ModelTemplate string `json:"model_template"`
	Namespace     string `json:"namespace"`
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
