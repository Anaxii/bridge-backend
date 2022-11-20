package wallet

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/config"
	"puffinbridgebackend/global"
)

func Balance(network global.Networks) *big.Int {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:", err)
		log.WithFields(log.Fields{
			"network":   network.Name,
			"err": err,
		}).Error("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:")
		return big.NewInt(0)
	}

	balance, err := conn.BalanceAt(context.Background(), common.HexToAddress(config.PublicKey), nil)
	if err != nil {
		log.WithFields(log.Fields{
			"network":   network.Name,
			"err": err,
		}).Error("wallet/wallet.go:Balance(): Failed to call balance")
		return big.NewInt(0)
	}

	return balance
}

func Block(network global.Networks) *big.Int {

	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:", err)
		log.WithFields(log.Fields{
			"network":   network.Name,
			"err": err,
		}).Error("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:")
		return big.NewInt(0)
	}

	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:", err)
		log.WithFields(log.Fields{
			"network":   network.Name,
			"err": err,
		}).Error("wallet/wallet.go:Balance(): Failed to connect to the Ethereum client:")
		return big.NewInt(0)
	}

	return header.Number
}

