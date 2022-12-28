package handler

import (
	"puffinbridgebackend/internal/api"
	"time"
)

func (h *Handler) writeLogs(v BridgeRequest, status string, err error) {
	_err := ""
	if err != nil {
		_err = err.Error()
	}

	go api.Log(LogHistory{Log: v, Status: status, Error: _err, Timestamp: time.Now().Unix()})
}
