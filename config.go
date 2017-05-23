package main

import (
	"encoding/json"
	"io/ioutil"
)

type mainConfig struct {
	Token string
	Channels struct {
		GreetingChannel string
	}
}

// Loads JSON config into supplied struct, optionally panics on error
// Config struct fields must be exported
func loadConfig(configname string, configstruct interface{}, panicerr bool) error {
	io, err := ioutil.ReadFile(configname + ".json")
	if err != nil {
		if panicerr {
			panic(err) // TODO: Change to proper logging once made
		}
		return err
	}
	
	err = json.Unmarshal(io, &configstruct)
	if err != nil {
		if panicerr {
			panic(err) // TODO: Change to proper logging once made
		}
		return err
	}
	
	return nil
}