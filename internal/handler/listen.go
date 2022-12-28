package handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/internal/api"
	"puffinbridgebackend/internal/blockchain"
	"puffinbridgebackend/internal/config"
	"puffinbridgebackend/internal/events"
	"puffinbridgebackend/pkg/db"
	"puffinbridgebackend/pkg/wallet"
	"time"
)

func (h *Handler) listenToEvents() {

	var event = make(chan config.NetworkLog)

	go h.handleQueue()

	go func() {
		ticker := time.NewTicker(15 * time.Second)
		x := 3
		for range ticker.C {
			for _, v := range config.NetworksMap {
				h.updateLastBlock(v, x)
			}
			h.updateLastBlock(config.Subnet, x)

			if x == 3 {
				x = 0
			} else {
				x++
			}
		}
	}()

	log.Info("Starting event listeners")
	for _, v := range config.NetworksMap {
		go blockchain.ListenToEvents(v, v.BridgeAddress, event)
	}
	go blockchain.ListenToEvents(config.Subnet, config.Subnet.BridgeAddress, event)

	for {
		select {
		case vLog := <-event:
			data, method, err := events.FindEvent([]types.Log{vLog.Log}, h.BridgeABI)
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Info("Unable to parse event")
			}
			h.handleEvent(data, method, vLog.Network)
			if len(h.BridgeQueue) == 0 {
				db.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))
			}
		}
	}
}

func (h *Handler) updateLastBlock(v config.Networks, x int) {
	walletBlock, _ := wallet.Block(v)
	h.Blocks[v.Name] = int(walletBlock.Int64())
	if walletBlock.Int64() > 0 {
		if len(h.BridgeQueue) == 0 {
			log.WithFields(log.Fields{
				"block":   walletBlock.Int64(),
				"network": v.Name,
			}).Info("Updated last synced block")

			if x == 3 {
				go api.Log(LogHistory{Status: "Updated last synced block", Message: "lastBlock", Log: BridgeRequest{Block: walletBlock.Int64(), NetworkIn: v}, Timestamp: time.Now().Unix()})
			}

			db.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", walletBlock.Int64())))
		} else {
			log.WithFields(log.Fields{
				"block":      walletBlock.Int64(),
				"queue_size": len(h.BridgeQueue),
				"network":    v.Name,
			}).Info("Last block")
			for _, v := range h.BridgeQueue {
				log.Info(v.Method)
			}
		}
	}
}
