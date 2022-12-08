package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"io"
	_log "log"
	"math/rand"
	"net/http"
	"puffinbridgebackend/config"
	"puffinbridgebackend/global"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 100,
	WriteBufferSize: 1024 * 100,
}

func startRest() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/logs", getLogs).Methods("POST")
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

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	data := make(chan interface{})
	id := RandStringRunes(20)
	global.SocketChannels[id] = data
	reader(ws, data, id)
}
