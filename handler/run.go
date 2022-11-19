package handler

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
)

func RunHandler() {
	listenToEvents()
}

func listenToEvents() {

	var event = make(chan types.Log)

	logBridgeInSig := []byte("BridgeIn(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)")
	logBridgeOutSig := []byte("BridgeOut(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)")
	logBridgeOutWarmSig := []byte("BridgOutWarm(address indexed user, address indexed asset, uint256 indexed amount, bytes32 id)")
	logBridgeOutCanceledSig := []byte("BridgOutCanceled(address indexed user, bytes32 id)")

	logBridgeInSigHash := crypto.Keccak256Hash(logBridgeInSig)
	logBridgeOutSigHash := crypto.Keccak256Hash(logBridgeOutSig)
	logBridgeOutWarmSigHash := crypto.Keccak256Hash(logBridgeOutWarmSig)
	logBridgeOutCanceledSigHash := crypto.Keccak256Hash(logBridgeOutCanceledSig)


	for _, v := range config.Networks {
		go contractInteraction.ListenToEvents(v, v.BridgeAddress, &event)
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
			switch vLog.Topics[0].Hex() {
			case logBridgeInSigHash.Hex():
				//
			case logBridgeOutSigHash.Hex():
				//
			case logBridgeOutWarmSigHash.Hex():
				//
			case logBridgeOutCanceledSigHash.Hex():
				//
			}
		}
	}
}
