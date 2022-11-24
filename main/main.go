package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/api"
	"puffinbridgebackend/global"
	"puffinbridgebackend/handler"
	"puffinbridgebackend/initialize"
	"puffinbridgebackend/state"
)

func main() {
	log.Info("Initializing handler")
	initialize.RunPreChecks()
	log.Info("Starting handler")

	go api.RunAPI()
	var _handler = handler.Handler{}

	var logHistory []handler.LogHistory
	logs, err := state.Read([]byte("logs"), []byte("logs"))

	err = json.Unmarshal(logs, &logHistory)
	if err == nil {
		_handler.Logs = logHistory
	}

	global.Logs = _handler.Logs

	_handler.RunHandler()
}
