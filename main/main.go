package main

import (
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/initialize"
)
func main() {
	log.Info("Initializing handler")
	initialize.RunPreChecks()
}

