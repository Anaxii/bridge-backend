package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/api"
	"puffinbridgebackend/config"
	"puffinbridgebackend/embeddeddatabase"
	"puffinbridgebackend/global"
	"puffinbridgebackend/handler"
	"puffinbridgebackend/initialize"
)

func main() {
	log.Info("Initializing handler")

	if config.APIEnabled {
		log.Info("Starting API and Websocket")
		go api.RunAPI()
	}

	initialize.RunPreChecks()
	log.Info("Starting handler")

	var _handler = handler.Handler{}

	var logHistory []handler.LogHistory
	logs, err := embeddeddatabase.Read([]byte("logs"), []byte("logs"))

	err = json.Unmarshal(logs, &logHistory)
	if err == nil {
		_handler.Logs = logHistory
	}

	for _, v := range _handler.Logs {
		global.Logs = append(global.Logs, v)
	}

	_handler.RunHandler()
}
