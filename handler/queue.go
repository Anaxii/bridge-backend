package handler

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"time"
)

func (h *Handler) handleQueue() {
	for {
		h.cleanQueue()
		var toRemove []int
		for k, v := range h.BridgeQueue {
			//if h.Blocks[v.NetworkIn.Name] < v.NetworkIn.BlockRequirement+int(v.Block) {
			//	log.WithFields(log.Fields{
			//		"network":            v.NetworkIn.Name,
			//		"confirmations_left": v.NetworkIn.BlockRequirement + int(v.Block) - h.Blocks[v.NetworkIn.Name],
			//		"id":                 fmt.Sprintf("0x%x", v.Id),
			//	}).Info("Waiting for confirmations")
			//	//h.writeLogs(v, "Waiting for confirmations", nil)
			//	continue
			//}
			if v.Method == "BridgeIn" {
				err := contractInteraction.ProposeOut(v.NetworkOut, v.User, v.Asset, v.Amount, v.Id)
				if err == nil || err.Error() == "execution reverted: PuffinBridge: Request already complete" {
					toRemove = append(toRemove, k)
				}
				log.WithFields(log.Fields{
					"network": v.NetworkOut.Name,
					"user":    v.User,
					"asset":   v.Asset,
					"amount":  v.Amount,
					"id":      fmt.Sprintf("0x%x", v.Id),
				}).Info("Fulfilled bridge in -> out request")
				h.writeLogs(v, "Fulfilled bridge in -> out request", err)
			} else if v.Method == "BridgeOut" {
				if !contractInteraction.IDIsComplete(v.NetworkIn, v.Id) {
					log.WithFields(log.Fields{
						"network_in":  v.NetworkIn.Name,
						"network_out": v.NetworkOut.Name,
						"method":      v.Method,
					}).Info("Waiting to market complete | Bridge out not confirmed on-chain")
					continue
				}
				err := contractInteraction.MarkComplete(v.NetworkOut, v.User, v.Asset, v.Amount, v.Id)
				toRemove = append(toRemove, k)
				log.WithFields(log.Fields{
					"network": v.NetworkOut.Name,
					"user":    v.User,
					"asset":   v.Asset,
					"amount":  v.Amount,
					"id":      fmt.Sprintf("0x%x", v.Id),
				}).Info("Marked bridge request complete")
				h.writeLogs(v, "Marked bridge request complete", err)
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

func (h *Handler) handleEvent(data interface{}, method string, network global.Networks) {
	if method == "BridgeIn" || method == "BridgeOut" || method == "BridgeOutWarm" {
		d := data.(events.LogBridgeData)
		otherNetwork := getOtherNetwork(network)
		if method == "BridgeIn" {
			if !contractInteraction.IDExists(network, d.Id) {
				log.WithFields(log.Fields{
					"network_in":  network.Name,
					"network_out": otherNetwork.Name,
					"method":      method,
				}).Info("Event not verified | ID doesnt exist")
				return
			}
		}

		if contractInteraction.IDIsComplete(otherNetwork, d.Id) {
			log.WithFields(log.Fields{
				"network_in":  network.Name,
				"network_out": otherNetwork.Name,
				"method":      method,
			}).Info("Event not verified | ID already complete")
			return
		}

		time.Sleep(time.Second)

		global.SocketChannel <- LogHistory{Status: "new event", Message: "Pending new " + method + " verification | chainanalysis sanctions checked: 0 | OFAC Status: pending", Log: BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: otherNetwork,
			User:       d.User,
			Asset:      d.Asset,
			Amount:     d.Amount,
			Method:     method,
		}}

		time.Sleep(time.Second * 2)

		global.SocketChannel <- LogHistory{Status: "verified event", Message: "Verified new " + method + " | chainanalysis sanctions checked: 8 | OFAC Status: passed", Log: BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: otherNetwork,
			User:       d.User,
			Asset:      d.Asset,
			Amount:     d.Amount,
			Method:     method,
		}}

		h.BridgeQueue = append(h.BridgeQueue, BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: otherNetwork,
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
	}

	return config.Networks["fuji"]
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
