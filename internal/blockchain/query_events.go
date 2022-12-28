package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/internal/config"
	"puffinbridgebackend/internal/events"
)

func QueryEvent(network config.Networks, blockStart int64, blockEnd int64, address string, abi abi.ABI) (interface{}, string, error) {
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
			"location": "blockchain/contractInteraction/query_events.go:QueryEvent:26",
		}).Fatal(err)
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Error(err)
		return nil, "", err
	}

	data, method, err := events.FindEvent(logs, abi)

	return data, method, err

}
