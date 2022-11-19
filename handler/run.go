package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/contractABI"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"strings"
)

func RunHandler() {
	listenToEvents()
}

func listenToEvents() {

	var event = make(chan global.NetworkLog)

	logBridgeInSig := []byte("BridgeIn(address,address,uint256,bytes32)")
	logBridgeOutSig := []byte("BridgeOut(address,address,uint256,bytes32)")
	logBridgeOutWarmSig := []byte("BridgeOutWarm(address,address,uint256,bytes32)")
	logBridgeOutCanceledSig := []byte("BridgeOutCanceled(address,bytes32)")

	logBridgeInSigHash := crypto.Keccak256Hash(logBridgeInSig)
	logBridgeOutSigHash := crypto.Keccak256Hash(logBridgeOutSig)
	logBridgeOutWarmSigHash := crypto.Keccak256Hash(logBridgeOutWarmSig)
	logBridgeOutCanceledSigHash := crypto.Keccak256Hash(logBridgeOutCanceledSig)

	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinMainnetBridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range config.Networks {
		go contractInteraction.ListenToEvents(v, v.BridgeAddress, event)
	}
	//block, err := state.Read([]byte("state"), []byte("block"))
	//lastBlock, err := strconv.Atoi(string(block))
	//if err != nil {
	//	log.Warn("Invalid block state, setting to 0")
	//	lastBlock = 0
	//}

	for {
		select {
		case vLog := <-event:
			switch vLog.Log.Topics[0].Hex() {
			case logBridgeInSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeIn", vLog.Log)
			case logBridgeOutSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeOut", vLog.Log)
			case logBridgeOutWarmSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeOutWarm", vLog.Log)
			case logBridgeOutCanceledSigHash.Hex():
				_, err = events.LogBridgeCancel(bridgeABI, "BridgeCancel", vLog.Log)
			}
			state.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))

		}
	}
}
