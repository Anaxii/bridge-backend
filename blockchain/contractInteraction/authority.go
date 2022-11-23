package contractInteraction

import (
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"puffinbridgebackend/config"
	abi "puffinbridgebackend/contractABI"
	"puffinbridgebackend/global"
)


func IsKYCOnMainnet(network global.Networks) bool {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return false
	}

	verify, err := abi.NewPuffinApprovals(common.HexToAddress(network.KYCAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return false
	}

	isApproved, err := verify.IsApproved(nil, common.HexToAddress(config.PublicKey))
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}
	return isApproved
}

func IsKYCOnSubnet(network global.Networks) bool {
	conn, err := ethclient.Dial(network.RpcURL)
	if err != nil {
		log.Println("Failed to connect to the Ethereum client:", err)
		return false
	}

	verifyPuffin, err := abi.NewAllowListInterface(common.HexToAddress(network.KYCAddress), conn)
	if err != nil {
		log.Println("Failed to instantiate PuffinApprovedAccounts contract:", err)
		return false
	}

	isEnabled, err := verifyPuffin.ReadAllowList(nil, common.HexToAddress(config.PublicKey))
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	return isEnabled != big.NewInt(0)
}

func IsVoterOnMainnet(network global.Networks) bool {
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

	isEnabled, err := verifyPuffin.IsVoter(nil, common.HexToAddress(config.PublicKey))
	if err != nil {
		log.Println("Failed to read user:", err)
		return false
	}

	return isEnabled
}



