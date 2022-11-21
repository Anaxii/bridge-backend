package events

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/global"
)

func FindEvent(logs []types.Log, abi ethABI.ABI) error {
	var err error
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case global.LogBridgeInSigHash.Hex():
			_, err = logBridge(abi, "BridgeIn", vLog)
		case global.LogBridgeOutSigHash.Hex():
			_, err = logBridge(abi, "BridgeOut", vLog)
		case global.LogBridgeOutWarmSigHash.Hex():
			_, err = logBridge(abi, "BridgeOutWarm", vLog)
		case global.LogBridgeOutCanceledSigHash.Hex():
			_, err = logBridgeCancel(abi, "BridgeCancel", vLog)

		}
	}
	return err
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
	return data, nil
}
