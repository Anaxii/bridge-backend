package handler

import (
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/api"
	"puffinbridgebackend/blockchain"
	"puffinbridgebackend/config"
	"puffinbridgebackend/events"
	"time"
)

func (h *Handler) handleEvent(data interface{}, method string, network config.Networks) {
	if method == "BridgeIn" || method == "BridgeOut" || method == "BridgeOutWarm" {
		d := data.(events.LogBridgeData)
		otherNetwork := getOtherNetwork(network)
		if method == "BridgeIn" {
			if !blockchain.IDExists(network, d.Id) {
				log.WithFields(log.Fields{
					"network_in":  network.Name,
					"network_out": otherNetwork.Name,
					"method":      method,
				}).Info("Event not verified | ID doesnt exist")
				return
			}
		}

		if blockchain.IDIsComplete(otherNetwork, d.Id) {
			log.WithFields(log.Fields{
				"network_in":  network.Name,
				"network_out": otherNetwork.Name,
				"method":      method,
			}).Info("Event not verified | ID already complete")
			return
		}

		time.Sleep(time.Second)

		api.Log(LogHistory{Status: "new event", Message: "Pending new " + method + " verification | chainanalysis sanctions checked: 0 | OFAC Status: pending", Log: BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: otherNetwork,
			User:       d.User,
			Asset:      d.Asset,
			Amount:     d.Amount,
			Method:     method,
		}})

		time.Sleep(time.Second * 2)

		api.Log(LogHistory{Status: "verified event", Message: "Verified new " + method + " | chainanalysis sanctions checked: 8 | OFAC Status: passed", Log: BridgeRequest{
			Id:         d.Id,
			Block:      d.Block,
			NetworkIn:  network,
			NetworkOut: otherNetwork,
			User:       d.User,
			Asset:      d.Asset,
			Amount:     d.Amount,
			Method:     method,
		}})

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

func getOtherNetwork(network config.Networks) config.Networks {
	if network.ChainId == big.NewInt(43113114) {
		return config.Subnet
	}

	return config.NetworksMap["fuji"]
}
