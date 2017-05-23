package main

import (
	"fmt"
)

func addCommand(name string, function interface{}) {
	logInfo(fmt.Sprintf("Adding command [%s]", name))
	commandList[name] = function
}
