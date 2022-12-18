package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"puffinbridgebackend/api"
	"puffinbridgebackend/config"
	"puffinbridgebackend/db"
	"puffinbridgebackend/handler"
	"puffinbridgebackend/initialize"
)

func main() {

	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	//defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})

	log.Info("Initializing handler")

	if config.APIEnabled {
		log.Info("Starting API and Websocket")
		go api.RunAPI()
	}

	initialize.RunPreChecks()
	log.Info("Starting handler")

	var _handler = handler.Handler{}

	var logHistory []handler.LogHistory
	logs, err := db.Read([]byte("logs"), []byte("logs"))

	err = json.Unmarshal(logs, &logHistory)


	_handler.RunHandler()
}
