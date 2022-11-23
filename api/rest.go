package api

import (
	"github.com/gorilla/mux"
	_log "log"
	"net/http"
	"puffinbridgebackend/config"
)

func startRest() {

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/logs", getLogs).Methods("POST")

	r.Use(mux.CORSMethodMiddleware(r))
	_log.Fatal(http.ListenAndServe(":" + config.APIPort, r))
}

func getLogs(w http.ResponseWriter, r *http.Request) {


	//w.Write()

}