package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Configuration for loading from JSON
type Configuration struct {
	Listen             string `json:"listen"`
	StaticFilesFolder  string `json:"staticFilesFolder"`
	CommandFilesFolder string `json:"commandFilesFolder"`
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
	if len(conf.Listen) == 0 {
		conf.Listen = ":8000"
	}
	if len(conf.StaticFilesFolder) == 0 {
		conf.StaticFilesFolder = "./static"
	}
	if len(conf.CommandFilesFolder) == 0 {
		conf.CommandFilesFolder = "./commands"
	}
	return conf
}
