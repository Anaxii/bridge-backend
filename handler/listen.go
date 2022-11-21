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

func (h Handler) listenToEvents() {

	var event = make(chan global.NetworkLog)

	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinMainnetBridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Syncing networks")
	numSynced := 0
	x := 0
	for numSynced < len(config.Networks) {
		numSynced = 0
		x++
		for _, v := range config.Networks {
			numSynced = sync(v, x, numSynced, bridgeABI)
		}
	}

	x = 0
	numSynced = 0
	for numSynced < 1 {
		x++
		numSynced = sync(config.Subnet, x, numSynced, bridgeABI)
	}

	go func() {
		ticker := time.NewTicker(15 * time.Second)

		for range ticker.C {
			for _, v := range config.Networks {
				updateLastBlock(v)
			}
			updateLastBlock(config.Subnet)
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
			events.FindEvent([]types.Log{vLog.Log}, bridgeABI)
			state.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))
		}
	}
}

func updateLastBlock(v global.Networks) {
	walletBlock := wallet.Block(v)
	if walletBlock.Int64() > 0 {
		log.WithFields(log.Fields{
			"block":   walletBlock.Int64(),
			"network": v.Name,
		}).Info("Updated last synced block")
		state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", walletBlock.Int64())))
	}
}
