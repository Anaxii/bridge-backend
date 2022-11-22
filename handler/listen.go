package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	abi "puffinbridgebackend/contractABI"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"puffinbridgebackend/wallet"
	"strings"
	"time"
)

func (h *Handler) listenToEvents() {

	var event = make(chan global.NetworkLog)

	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinMainnetBridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Syncing networks")
	startBlock := map[string]int{}
	numSynced := 0
	for x := 0; numSynced < len(config.Networks); x++ {
		numSynced = 0
		for _, v := range config.Networks {
			_numSynced, lastBlock := h.sync(v, x, numSynced, bridgeABI)
			if x == 0 {
				startBlock[v.Name] = lastBlock
			}
			numSynced = _numSynced
		}
	}

	numSynced = 0
	for x := 0; numSynced < 1; x++ {
		_numSynced, lastBlock := h.sync(config.Subnet, x, numSynced, bridgeABI)
		if x == 0 {
			startBlock[config.Subnet.Name] = lastBlock
		}
		numSynced = _numSynced
	}

	for k, v := range startBlock {
		state.Write([]byte("block"), []byte(k), []byte(fmt.Sprintf("%v", v)))
	}

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
			data, method, err := events.FindEvent([]types.Log{vLog.Log}, bridgeABI)
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
		if len(h.BridgeQueue) == 0 {
			log.WithFields(log.Fields{
				"block":   walletBlock.Int64(),
				"queue_size": len(h.BridgeQueue),
				"network": v.Name,
			}).Info("Updated last synced block")
			state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", walletBlock.Int64())))
		}
	}
}
