package contractInteraction

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/global"
)

func ListenToEvents(network global.Networks, _contractAddress string, events *chan types.Log) {
	client, err := ethclient.Dial(network.WSURL)
	if err != nil {
		log.WithFields(log.Fields{
			"network":   network.Name,
			"contract": _contractAddress,
			"location": "blockchain/contractInteraction/listen.go:ListenToEvents:19",
		}).Fatal(err)
	}

	contractAddress := common.HexToAddress(_contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, *events)
	if err != nil {
		log.WithFields(log.Fields{
			"network":   network.Name,
			"contract": _contractAddress,
			"location": "blockchain/contractInteraction/listen.go:ListenToEvents:33",
		}).Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.WithFields(log.Fields{
				"network":   network.Name,
				"contract": _contractAddress,
				"location": "blockchain/contractInteraction/listen.go:ListenToEvents:44",
			}).Error("Event listener died, rebooting |", err)
			ListenToEvents(network, _contractAddress, events)
			return
		}
	}

}
