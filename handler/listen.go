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
	"strings"
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
