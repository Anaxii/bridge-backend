package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
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

	defer r.Body.Close()

	if len(global.Logs) == 0 {
		w.WriteHeader(http.StatusTooEarly)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var body map[string]int

	err = json.Unmarshal(b, &body)
	if err != nil {
		log.Println(err, string(b))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lines := 50

	if _, ok := body["lines"]; ok {
		lines = body["lines"]
	}

	if lines > len(global.Logs) {
		lines = len(global.Logs)
	}

	data, err := json.Marshal(global.Logs[:len(global.Logs)-1])
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

	reader(ws)
}
