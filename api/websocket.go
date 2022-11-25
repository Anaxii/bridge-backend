package api

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/global"
)

func reader(conn *websocket.Conn) {
	var messageType int
	enabled := map[string]bool{"logs": false}
	go func() {
		for {
			// read in a message
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			if err := conn.WriteMessage(messageType, []byte("Connection to Puffin Bridge handler established")); err != nil {
				log.Println(err)
				return
			}

			if string(msg) == "logs" {
				enabled["logs"] = true
				if err := conn.WriteMessage(messageType, []byte("Logs enabled")); err != nil {
					log.Println(err)
					return
				}
			}

		}
	}()
	for {
		select {
		case d := <-global.SocketChannel:
			log.Println("CHECK")
			if err := conn.WriteMessage(messageType, d.([]byte)); err != nil {
				return
			}
		}
	}

}
