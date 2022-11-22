package events

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/global"
)

func FindEvent(logs []types.Log, abi ethABI.ABI) (interface{}, string, error) {
	var err error
	var data interface{}
	method := ""
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case global.LogBridgeInSigHash.Hex():
			data, err = logBridge(abi, "BridgeIn", vLog)
			method = "BridgeIn"
		case global.LogBridgeOutSigHash.Hex():
			data, err = logBridge(abi, "BridgeOut", vLog)
			method = "BridgeOut"
		case global.LogBridgeOutWarmSigHash.Hex():
			data, err = logBridge(abi, "BridgeOutWarm", vLog)
			method = "BridgeOutWarm"
		case global.LogBridgeOutCanceledSigHash.Hex():
			data, err = logBridgeCancel(abi, "BridgeCancel", vLog)
			method = "BridgeCancel"

		}
	}
	return data, method, err
}

func logBridge(_abi ethABI.ABI, method string, vLog types.Log) (LogBridgeData, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":   vLog.BlockNumber,
		"tx_hash": vLog.TxHash,
		"method": method,
	}).Info("New bridge event")

	var data LogBridgeData
	err := _abi.UnpackIntoInterface(&data, method, vLog.Data)
	if err != nil {
		log.WithFields(log.Fields{
			"bridge_address": vLog.Address,
			"block":   vLog.BlockNumber,
			"tx_hash": vLog.TxHash,
			"method": method,
		}).Error("Could not unpack event")
		return LogBridgeData{}, err
	}
	data.User = common.HexToAddress(vLog.Topics[1].Hex())
	data.Asset = common.HexToAddress(vLog.Topics[2].Hex())
	data.Amount = vLog.Topics[3].Big()
	data.Block = int64(vLog.BlockNumber)

	return data, nil
}

func logBridgeCancel(_abi ethABI.ABI, method string, vLog types.Log) (LogBridgeCancelData, error) {
	log.WithFields(log.Fields{
		"bridge_address": vLog.Address,
		"block":   vLog.BlockNumber,
		"tx_hash": vLog.TxHash,
		"method": method,
	}).Info("New bridge event")

	var data LogBridgeCancelData
	if len(vLog.Topics) == 3 {
		data.User = common.HexToAddress(vLog.Topics[1].Hex())
		data.Id = common.HexToHash(vLog.Topics[2].Hex())
	}
	data.Block = int64(vLog.BlockNumber)
	return data, nil
}
