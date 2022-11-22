package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/config"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"time"
)

func (h *Handler) handleQueue() {
	for {
		h.cleanQueue()
		for _, v := range h.BridgeQueue {
			if v.Method == "BridgeIn" {

			}
			log.Println(v)
		}
		time.Sleep(time.Second * 15)
	}
}

func (h *Handler) handleEvent(data interface{}, method string, network global.Networks) {
	if method == "BridgeIn" || method == "BridgeOut" || method == "BridgeOutWarm" {
		d := data.(events.LogBridgeData)
		if !Verify(network, getOtherNetwork(network), d.Id) {
			return
		}
		h.BridgeQueue = append(h.BridgeQueue, BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: getOtherNetwork(network),
			User:       d.User,
			Asset:      d.Asset,
			Amount:     d.Amount,
			Method:     method,
		})
	}
}

func getOtherNetwork(network global.Networks) global.Networks {
	if network.Name == "fuji" {
		return config.Subnet
	} else {
		return config.Networks["fuji"]
	}
}

func (h *Handler) cleanQueue() {
	isComplete := map[string]bool{}
	log.Println(len(h.BridgeQueue))
	for k, v := range h.BridgeQueue {
		if isComplete[v.NetworkOut.Name + fmt.Sprintf("%v", v.Id)] {
			if k < 0 || k >= len(h.BridgeQueue) {
				continue
			} else {
				h.BridgeQueue = append(h.BridgeQueue[:k], h.BridgeQueue[k+1:]...)
			}
		} else {
			isComplete[v.NetworkOut.Name + fmt.Sprintf("%v", v.Id)] = true
		}
	}
}
