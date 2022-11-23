package handler

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/config"
	abi "puffinbridgebackend/contractABI"
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
	for _, v := range config.Networks {
		walletBlock := wallet.Block(v)
		if walletBlock.Int64() > 0 {
			h.Blocks[v.Name] = int(walletBlock.Int64())
		}
	}
	walletBlock := wallet.Block(config.Subnet)
	if walletBlock.Int64() > 0 {
		h.Blocks[config.Subnet.Name] = int(walletBlock.Int64())
	}

	h.startSync()
	h.listenToEvents()
}
