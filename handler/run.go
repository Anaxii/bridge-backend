package handler

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	abi "puffinbridgebackend/blockchain/contractABI"
	"puffinbridgebackend/config"
	"puffinbridgebackend/wallet"
	"strings"
)

func (h *Handler) RunHandler() {
	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinBridgeABI))
	if err != nil {
		log.Fatal(err)
	}
	h.BridgeABI = bridgeABI

	h.Blocks = map[string]int{}
	for _, v := range config.NetworksMap {
		walletBlock, _ := wallet.Block(v)
		if walletBlock.Int64() > 0 {
			h.Blocks[v.Name] = int(walletBlock.Int64())
		}
	}

	walletBlock, _ := wallet.Block(config.Subnet)
	if walletBlock.Int64() > 0 {
		h.Blocks[config.Subnet.Name] = int(walletBlock.Int64())
	}

	h.startSync()
	h.listenToEvents()
}
