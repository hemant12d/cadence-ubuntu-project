package main

import (
	"github.com/fatih/color"
)

var HostPort = "127.0.0.1:7933"
var Domain = "SimpleDomain"
var TaskListName = "my-first-worker"
var ClientName = "my-first-client"
var CadenceService = "cadence-frontend"

func main() {
	color.Cyan("Let's get started")
	StartWorker(BuildLogger(), BuildCadenceClient())
}
