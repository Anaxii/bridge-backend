package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	_log "log"
	"net/http"
	"puffinbridgebackend/config"
	"puffinbridgebackend/global"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func startRest() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/logs", getLogs).Methods("POST")
	r.HandleFunc("/ws", getWS).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	_log.Fatal(http.ListenAndServe(":"+config.APIPort, r))
}

func getLogs(w http.ResponseWriter, r *http.Request) {

	data, err := json.Marshal(global.Logs)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(string(data))
	w.Write(data)

}

func getWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	reader(ws)
}
