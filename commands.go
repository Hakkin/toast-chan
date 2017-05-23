package main

func addCommand(name string, function interface{}) {
	commandList[name] = function
}