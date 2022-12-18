package blockchain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/config"
)

func ListenToEvents(network config.Networks, _contractAddress string, events chan config.NetworkLog) {
	client, err := ethclient.Dial(network.WSURL)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"contract": _contractAddress,
			"location": "blockchain/contractInteraction/listen_to_events.go:ListenToEvents:19",
		}).Fatal(err)
	}

	contractAddress := common.HexToAddress(_contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var event = make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, event)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"contract": _contractAddress,
			"location": "blockchain/contractInteraction/listen_to_events.go:ListenToEvents:33",
		}).Fatal(err)
	}
	go func() {
		for {
			select {
			case ev := <-event:
				events <- config.NetworkLog{Network: network, Log: ev}
			}
		}
	}()

	for {
		select {
		case err := <-sub.Err():
			log.WithFields(log.Fields{
				"network":  network.Name,
				"contract": _contractAddress,
				"location": "blockchain/contractInteraction/listen_to_events.go:ListenToEvents:44",
			}).Error("Event listener died, rebooting |", err)
			go ListenToEvents(network, _contractAddress, events)
			return
		}
	}

}
