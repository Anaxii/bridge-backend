package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/internal/blockchain"
)

func (h *Handler) bridgeIn(v BridgeRequest) bool {
	err := blockchain.ProposeOut(v.NetworkOut, v.User, v.Asset, v.Amount, v.Id)
	if err == nil || err.Error() == "execution reverted: PuffinBridge: Request already complete" {

		return true
	}
	log.WithFields(log.Fields{
		"network": v.NetworkOut.Name,
		"user":    v.User,
		"asset":   v.Asset,
		"amount":  v.Amount,
		"id":      fmt.Sprintf("0x%x", v.Id),
	}).Info("Fulfilled bridge in -> out request")
	h.writeLogs(v, "Fulfilled bridge in -> out request", err)
	return false
}

func (h *Handler) bridgeOut(v BridgeRequest) bool {
	if !blockchain.IDIsComplete(v.NetworkIn, v.Id) {
		log.WithFields(log.Fields{
			"network_in":  v.NetworkIn.Name,
			"network_out": v.NetworkOut.Name,
			"method":      v.Method,
		}).Info("Waiting to market complete | Bridge out not confirmed on-chain")
		return false
	}
	err := blockchain.MarkComplete(v.NetworkOut, v.User, v.Asset, v.Amount, v.Id)
	log.WithFields(log.Fields{
		"network": v.NetworkOut.Name,
		"user":    v.User,
		"asset":   v.Asset,
		"amount":  v.Amount,
		"id":      fmt.Sprintf("0x%x", v.Id),
	}).Info("Marked bridge request complete")
	h.writeLogs(v, "Marked bridge request complete", err)
	return true
}
