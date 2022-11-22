package handler

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	abi "puffinbridgebackend/contractABI"
	"strings"
)

func (h *Handler) RunHandler() {
	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinMainnetBridgeABI))
	if err != nil {
		log.Fatal(err)
	}
	h.BridgeABI = bridgeABI
	h.startSync()
	h.listenToEvents()
}
