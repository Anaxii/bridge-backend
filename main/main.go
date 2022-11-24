package main

import (
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/api"
	"puffinbridgebackend/handler"
	"puffinbridgebackend/initialize"
)

func main() {
	log.Info("Initializing handler")
	initialize.RunPreChecks()
	log.Info("Starting handler")

	go api.RunAPI()
	var _handler = handler.Handler{}
	_handler.RunHandler()
}
