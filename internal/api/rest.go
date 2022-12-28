package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	_log "log"
	"net/http"
	"puffinbridgebackend/internal/config"
	"puffinbridgebackend/pkg/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 100,
	WriteBufferSize: 1024 * 100,
}

func startRest() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/pub", getPub).Methods("GET")
	r.HandleFunc("/ws", getWS).Methods("GET")

	r.Use(mux.CORSMethodMiddleware(r))
	log.Info("API and websocket opened on port " + config.APIPort)
	_log.Fatal(http.ListenAndServe(":"+config.APIPort, r))
}

func getPub(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	data, err := json.Marshal(map[string]string{"pub": config.PublicKey})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(data)

}

func getWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	data := make(chan interface{})
	id := util.RandStringRunes(20)
	socketChannels[id] = data
	reader(ws, data, id)
}
