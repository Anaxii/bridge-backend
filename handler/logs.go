package handler

import (
	"encoding/json"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
)

func (h *Handler) writeLogs(v BridgeRequest, status string) {
	h.Logs = append(h.Logs, LogHistory{Log: v, Status: status})
	global.Logs = h.Logs
	data, _ := json.Marshal(h.Logs)
	global.SocketChannel <- data
	state.Write([]byte("logs"), []byte("logs"), data)
}
