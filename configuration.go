package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Configuration for loading from JSON
type Configuration struct {
	Commands []*Command `json:"commands"`
}

// LoadConfiguration from config.json file
func LoadConfiguration() *Configuration {
	var conf = &Configuration{}
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		panic(e)
	}
	err := json.NewDecoder(bytes.NewReader(file)).Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
