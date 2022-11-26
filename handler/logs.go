package handler

import (
	"encoding/json"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
)

func (h *Handler) writeLogs(v BridgeRequest, status string, err error) {
	h.Logs = append(h.Logs, LogHistory{Log: v, Status: status})
	global.Logs = append(global.Logs, LogHistory{Log: v, Status: status, Error: err})
	data, _ := json.Marshal(h.Logs)
	global.SocketChannel <- h.Logs[len(h.Logs)-1]
	state.Write([]byte("logs"), []byte("logs"), data)
}
