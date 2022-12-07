package handler

import (
	"encoding/json"
	"puffinbridgebackend/embeddeddatabase"
	"puffinbridgebackend/global"
	"time"
)

func (h *Handler) writeLogs(v BridgeRequest, status string, err error) {
	_err := ""
	if err != nil {
		_err = err.Error()
	}

	h.Logs = append(h.Logs, LogHistory{Log: v, Status: status, Error: _err, Timestamp: time.Now().Unix()})
	global.Logs = append(global.Logs, LogHistory{Log: v, Status: status, Error: _err, Timestamp: time.Now().Unix()})

	go global.Log(h.Logs[len(h.Logs)-1])
	data, _ := json.Marshal(h.Logs)
	embeddeddatabase.Write([]byte("logs"), []byte("logs"), data)
}
