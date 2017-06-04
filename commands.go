package main

import (
	"fmt"
)

var commandList = make(map[string]interface{})

func addCommand(name string, function interface{}) {
	logInfo(fmt.Sprintf("Adding command [%s]", name))
	commandList[name] = function
}
