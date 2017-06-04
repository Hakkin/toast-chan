package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const configDir = "./configs/"

// Loads JSON config into supplied struct, optionally panics on error
// Config struct fields must be exported
func loadConfig(configname string, configstruct interface{}, panicerr bool) error {
	logInfo(fmt.Sprintf("Loading config [%s.json]", configname))
	io, err := ioutil.ReadFile(configDir + configname + ".json")
	if err != nil {
		if panicerr {
			logFatal(fmt.Sprintf("There was an error while loading config [%s.json]: %s", configname, err))
		}
		return err
	}

	err = json.Unmarshal(io, &configstruct)
	if err != nil {
		if panicerr {
			logFatal(fmt.Sprintf("There was an error while parsing config [%s.json]: %s", configname, err))
		}
		return err
	}

	return nil
}
