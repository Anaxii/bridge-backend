package handler

import "github.com/ethereum/go-ethereum/common"

type Handler struct {
	//EventTopicHashes EventTopicHashes
}

type EventTopicHashes struct {
	LogBridgeInSigHash          common.Hash
	LogBridgeOutSigHash         common.Hash
	LogBridgeOutWarmSigHash     common.Hash
	LogBridgeOutCanceledSigHash common.Hash
}

func (h Handler) RunHandler() {
	h.listenToEvents()
}
