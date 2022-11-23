package contractInteraction

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	abi "puffinbridgebackend/contractABI"
	"puffinbridgebackend/global"
)

func IDExists(network global.Networks, requestId [32]byte) bool {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return false
	}

	verifyPuffin, err := abi.NewPuffinBridge(common.HexToAddress(network.BridgeAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return false
	}

	exists, err := verifyPuffin.BridgeIds(nil, requestId)
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	return exists
}

func IDIsComplete(network global.Networks, requestId [32]byte) bool {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return false
	}

	verifyPuffin, err := abi.NewPuffinBridge(common.HexToAddress(network.BridgeAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return false
	}

	complete, err := verifyPuffin.BridgeOutComplete(nil, requestId)
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	return complete
}