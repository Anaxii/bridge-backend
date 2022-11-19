package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/contractABI"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"strings"
)

func RunHandler() {
	listenToEvents()
}

type LogBridge struct {
	User   common.Address
	Asset  common.Address
	Amount *big.Int
	Id     [32]byte
}

type LogBridgeCancel struct {
	User common.Address
	Id   [32]byte
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
				_, err = logBridge(bridgeABI, "BridgeIn", vLog.Log)
			case logBridgeOutSigHash.Hex():
				_, err = logBridge(bridgeABI, "BridgeOut", vLog.Log)
			case logBridgeOutWarmSigHash.Hex():
				_, err = logBridge(bridgeABI, "BridgeOutWarm", vLog.Log)
			case logBridgeOutCanceledSigHash.Hex():
				_, err = logBridgeCancel(bridgeABI, "BridgeCancel", vLog.Log)
			}
			state.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))

		}
	}
}

func logBridge(_abi ethABI.ABI, method string, vLog types.Log) (LogBridge, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":   vLog.BlockNumber,
		"tx_hash": vLog.TxHash,
		"method": method,
	}).Info("New bridge event")

	var data LogBridge
	err := _abi.UnpackIntoInterface(&data, method, vLog.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"bridge_address": vLog.Address,
			"block":   vLog.BlockNumber,
			"tx_hash": vLog.TxHash,
			"method": method,
		}).Error("Could not unpack event")
		return LogBridge{}, err
	}
	data.User = common.HexToAddress(vLog.Topics[1].Hex())
	data.Asset = common.HexToAddress(vLog.Topics[2].Hex())
	data.Amount = vLog.Topics[3].Big()

	return data, nil
}

func logBridgeCancel(_abi ethABI.ABI, method string, vLog types.Log) (LogBridgeCancel, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":   vLog.BlockNumber,
		"tx_hash": vLog.TxHash,
		"method": method,
	}).Info("New bridge event")

	var data LogBridgeCancel
	err := _abi.UnpackIntoInterface(&data, method, vLog.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"bridge_address": vLog.Address,
			"block":   vLog.BlockNumber,
			"tx_hash": vLog.TxHash,
			"method": method,
		}).Error("Could not unpack event")
		return LogBridgeCancel{}, err
	}
	data.User = common.HexToAddress(vLog.Topics[1].Hex())
	// Debug ID not marshaling
	log.Println(vLog.Topics)

	return data, nil
}
