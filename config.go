package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type mainConfig struct {
	Token    string
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
			logFatal(fmt.Sprintf("There was an error while loading config %s.json: %s", configname, err))
		}
		return err
	}

	err = json.Unmarshal(io, &configstruct)
	if err != nil {
		if panicerr {
			logFatal(fmt.Sprintf("There was an error while parsing config %s.json: %s", configname, err))
		}
		return err
	}

	return nil
}
