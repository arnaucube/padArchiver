package main

import (
	"encoding/json"
	"io/ioutil"
)

//Config is the configuration of the API server
type Config struct {
	Port string `json:"port"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
