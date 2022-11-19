package main

import (
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/handler"
	"puffinbridgebackend/initialize"
)
func main() {
	log.Info("Initializing handler")
	initialize.RunPreChecks()
	log.Info("Starting handler")
	handler.RunHandler()
}

