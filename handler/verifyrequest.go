package handler

import (
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/global"
)

func Verify(networkIn global.Networks, networkOut global.Networks, requestId [32]byte) bool {
	if !contractInteraction.IDExists(networkIn, requestId) {
		return false
	}
	if contractInteraction.IDIsComplete(networkOut, requestId) {
		return false
	}
	return true
}

