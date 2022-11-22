package handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"puffinbridgebackend/wallet"
	"time"
)

func (h *Handler) listenToEvents() {

	var event = make(chan global.NetworkLog)

	go h.handleQueue()

	go func() {
		ticker := time.NewTicker(15 * time.Second)

		for range ticker.C {
			for _, v := range config.Networks {
				h.updateLastBlock(v)
			}
			h.updateLastBlock(config.Subnet)
		}
	}()

	log.Info("Starting event listeners")
	for _, v := range config.Networks {
		go contractInteraction.ListenToEvents(v, v.BridgeAddress, event)
	}
	go contractInteraction.ListenToEvents(config.Subnet, config.Subnet.BridgeAddress, event)

	for {
		select {
		case vLog := <-event:
			data, method, err := events.FindEvent([]types.Log{vLog.Log}, h.BridgeABI)
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Info("Unable to parse event")
			}
			h.handleEvent(data, method, vLog.Network)
			if len(h.BridgeQueue) == 0 {
				state.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))
			}
		}
	}
}

func (h *Handler) updateLastBlock(v global.Networks) {
	walletBlock := wallet.Block(v)
	if walletBlock.Int64() > 0 {
		h.Blocks[v.Name] = int(walletBlock.Int64())
		if len(h.BridgeQueue) == 0 {
			log.WithFields(log.Fields{
				"block":      walletBlock.Int64(),
				"network":    v.Name,
			}).Info("Updated last synced block")
			state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", walletBlock.Int64())))
		} else {
			log.WithFields(log.Fields{
				"block":      walletBlock.Int64(),
				"queue_size": len(h.BridgeQueue),
				"network":    v.Name,
			}).Info("Last block")
		}
	}
}
