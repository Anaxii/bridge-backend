package contractInteraction

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
)

func QueryEvent(network global.Networks, blockStart int64, blockEnd int64, address string, abi abi.ABI) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockStart),
		ToBlock:   big.NewInt(blockEnd),
		Addresses: []common.Address{
			common.HexToAddress(address),
		},
	}

	client, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"contract": address,
			"location": "blockchain/contractInteraction/queryevents.go:QueryEvent:26",
		}).Fatal(err)
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case global.LogBridgeInSigHash.Hex():
			_, err = events.LogBridge(abi, "BridgeIn", vLog)
		case global.LogBridgeOutSigHash.Hex():
			_, err = events.LogBridge(abi, "BridgeOut", vLog)
		case global.LogBridgeOutWarmSigHash.Hex():
			_, err = events.LogBridge(abi, "BridgeOutWarm", vLog)
		case global.LogBridgeOutCanceledSigHash.Hex():
			_, err = events.LogBridgeCancel(abi, "BridgeCancel", vLog)

		}
	}


}