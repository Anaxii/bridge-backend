package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	_log "log"
	"net/http"
	"puffinbridgebackend/config"
	"puffinbridgebackend/global"
)

func startRest() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/logs", getLogs).Methods("POST")

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
