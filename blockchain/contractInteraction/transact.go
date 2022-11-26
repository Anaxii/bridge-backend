package contractInteraction

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/config"
	abi "puffinbridgebackend/contractABI"
	"puffinbridgebackend/global"
)

func ProposeOut(network global.Networks, user common.Address, asset common.Address, amount *big.Int, id [32]byte) error {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "ProposeOut",
		}).Warn("Failed to connect to the Ethereum client:", err)
		return err
	}

	bridge, err := abi.NewPuffinBridge(common.HexToAddress(network.BridgeAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "ProposeOut",
		}).Warn("Failed to instantiate PuffinMainnetBridge contract:", err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, network.ChainId)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "ProposeOut",
		}).Warn("Failed to create authorized transactor:", err)
		log.Println("Failed to create authorized transactor:", err)
		return err
	}

	_, err = bridge.ProposeOut(auth, asset, user, amount, id, network.ChainId)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "ProposeOut",
		}).Warn("Failed to send ProposeOut ", err)
		return err
	}
	return nil
}

func MarkComplete(network global.Networks, user common.Address, asset common.Address, amount *big.Int, id [32]byte) error {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "MarkComplete",
		}).Warn("Failed to connect to the Ethereum client:", err)
		return err
	}

	bridge, err := abi.NewPuffinBridge(common.HexToAddress(network.BridgeAddress), conn)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "MarkComplete",
		}).Warn("Failed to instantiate PuffinMainnetBridge contract:", err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, network.ChainId)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"id":       id,
			"function": "MarkComplete",
		}).Warn("Failed to create authorized transactor:", err)
		log.Println("Failed to create authorized transactor:", err)
		return err
	}

	_, err = bridge.MarkInComplete(auth, id)
	if err != nil {
		log.WithFields(log.Fields{
			"network":  network.Name,
			"user":     user,
			"asset":    asset,
			"amount":   amount,
			"id":       id,
			"function": "MarkComplete",
		}).Warn("Failed to send MarkInComplete", err)
		return err
	}

	return nil
}
