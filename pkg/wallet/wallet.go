package wallet

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/internal/config"
)

// DialEthClient attempts to connect to an Ethereum client using the given RPC URL
func DialEthClient(rpcURL string) (*ethclient.Client, error) {
	conn, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.WithFields(log.Fields{
			"rpcURL": rpcURL,
			"err":    err,
		}).Error("Failed to connect to the Ethereum client")
		return nil, err
	}

	return conn, nil
}

// Balance returns the balance of the wallet with the given public key on the given network
func Balance(network config.Networks) (*big.Int, error) {
	conn, err := DialEthClient(network.RpcURL)
	if err != nil {
		return big.NewInt(0), err
	}

	balance, err := conn.BalanceAt(context.Background(), common.HexToAddress(config.PublicKey), nil)
	if err != nil {
		log.WithFields(log.Fields{
			"network": network.Name,
			"err":     err,
		}).Error("Failed to call balance")
		return big.NewInt(0), err
	}

	return balance, nil
}

// Block returns the latest block number on the given network
func Block(network config.Networks) (*big.Int, error) {
	conn, err := DialEthClient(network.RpcURL)
	if err != nil {
		return big.NewInt(0), err
	}

	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.WithFields(log.Fields{
			"network": network.Name,
			"err":     err,
		}).Error("Failed to get block header")
		return big.NewInt(0), err
	}

	return header.Number, nil
}
