package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Port string `json:"port"`
}

func ReadConfiguration() (*Configuration, error) {
	file, err := os.Open("../config/conf.json")
	if err != nil {
		locallog.Error("error reading config %s", err.Error())
		return nil, err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		locallog.Error("error reading config %s", err.Error())
		return nil, err
	}

	return &configuration, nil
}
