package api

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func reader(conn *websocket.Conn, dataChannel chan interface{}, id string) {
	enabled := map[string]bool{"logs": false}
	x := 0
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				delete(socketChannels, id)
				log.Println(err)
				return
			}
			response := map[string]string{"status": "Connection to Puffin Bridge handler established"}
			data, _ := json.Marshal(response)
			if x == 0 {
				x++
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					delete(socketChannels, id)
					return
				}
			}

			if string(msg) == "logs" {
				enabled["logs"] = true
				response = map[string]string{"status": "Logs enabled"}
				data, _ = json.Marshal(response)
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					return
				}
			}
			if string(msg) == "ping" {
				response = map[string]string{"status": "pong"}
				data, _ = json.Marshal(response)
				if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
					log.Println(err)
					delete(socketChannels, id)
					return
				}
			}

		}
	}()
	for {
		select {
		case d := <-dataChannel:
			response := map[string]interface{}{"status": "log", "data": d}
			data, err := json.Marshal(response)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
				delete(socketChannels, id)
				return
			}
		}
	}

}
