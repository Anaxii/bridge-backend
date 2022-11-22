package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"puffinbridgebackend/global"
)

type Handler struct {
	BridgeQueue []BridgeRequest
	//EventTopicHashes EventTopicHashes
}

type BridgeRequest struct {
	Id         [32]byte
	NetworkIn  global.Networks
	NetworkOut global.Networks
	User       common.Address
	Asset      common.Address
	Amount     *big.Int
	Block      int64
	Method     string
}

type EventTopicHashes struct {
	LogBridgeInSigHash          common.Hash
	LogBridgeOutSigHash         common.Hash
	LogBridgeOutWarmSigHash     common.Hash
	LogBridgeOutCanceledSigHash common.Hash
}

func (h *Handler) RunHandler() {
	h.listenToEvents()
}
