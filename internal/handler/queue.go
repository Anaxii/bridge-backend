package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

func (h *Handler) handleQueue() {
	for {
		h.cleanQueue()
		var toRemove []int
		for k, v := range h.BridgeQueue {
			if h.Blocks[v.NetworkIn.Name] < v.NetworkIn.BlockRequirement+int(v.Block) {
				log.WithFields(log.Fields{
					"network":            v.NetworkIn.Name,
					"confirmations_left": v.NetworkIn.BlockRequirement + int(v.Block) - h.Blocks[v.NetworkIn.Name],
					"id":                 fmt.Sprintf("0x%x", v.Id),
				}).Info("Waiting for confirmations")
				h.writeLogs(v, "Waiting for confirmations", nil)
				continue
			}
			if v.Method == "BridgeIn" {
				if shouldRemove := h.bridgeIn(v); shouldRemove {
					toRemove = append(toRemove, k)
				}
			} else if v.Method == "BridgeOut" {
				if shouldRemove := h.bridgeOut(v); shouldRemove {
					toRemove = append(toRemove, k)
				}
			} else if v.Method == "BridgeOutWarm" {
				// not currently handling warm wallets
				toRemove = append(toRemove, k)
			}
		}
		for k, v := range toRemove {
			if v-k >= 0 {
				h.removeFromQueue(v - k)
			}
		}
		toRemove = []int{}
		time.Sleep(time.Second * 15)
	}
}

func (h *Handler) cleanQueue() {
	isComplete := map[string]bool{}

	for k, v := range h.BridgeQueue {
		key := fmt.Sprintf("%s%s%v", v.NetworkOut.Name, v.Method, v.Id)
		if isComplete[key] {
			h.BridgeQueue = append(h.BridgeQueue[:k], h.BridgeQueue[k+1:]...)
		} else {
			isComplete[key] = true
		}
	}
}

func (h *Handler) removeFromQueue(k int) {
	if k < 0 || k >= len(h.BridgeQueue) {
		log.Error("BRIDGE QUEUE PARSE SKIP", k, " ", len(h.BridgeQueue))
		return
	} else {
		h.BridgeQueue = append(h.BridgeQueue[:k], h.BridgeQueue[k+1:]...)
	}
}
