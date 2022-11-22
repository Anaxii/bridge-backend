package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"puffinbridgebackend/wallet"
	"strconv"
)

func (h *Handler) sync(v global.Networks, x int, numSynced int, abi ethABI.ABI) (int, int) {
	block, err := state.Read([]byte("block"), []byte(v.Name))
	walletBlock := wallet.Block(v)
	if err != nil {
		log.WithFields(log.Fields{
			"status": "syncing",
			"err":    err,
		}).Error("Unable to fetch block")
		block = walletBlock.Bytes()
	}

	lastBlock, err := strconv.Atoi(string(block))
	if err != nil {
		log.WithFields(log.Fields{
			"status": "syncing",
			"block":  string(block),
		}).Warn("Could not convert block number to int")
		lastBlock = int(walletBlock.Int64())
	}
	if x%50 == 0 || x == 0 {
		log.WithFields(log.Fields{
			"status":  "syncing",
			"block":   fmt.Sprintf("%v/%v", lastBlock, walletBlock.Int64()),
			"network": v.Name,
		}).Info("Syncing")
	}

	if big.NewInt(int64(lastBlock)).Cmp(walletBlock) >= 0 {
		numSynced++
		state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", int64(lastBlock))))
	} else {
		nextBlock := int64(lastBlock) + 10
		if nextBlock > walletBlock.Int64() {
			nextBlock = walletBlock.Int64()
		}
		data, method, err := contractInteraction.QueryEvent(v, int64(lastBlock), nextBlock, v.BridgeAddress, abi)
		h.handleEvent(data, method, v)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Info("Unable to parse event in sync")
		}
		err = state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", nextBlock)))
		if err != nil {
			log.WithFields(log.Fields{
				"status": "syncing",
				"block":  string(block),
			}).Warn("Could not update database")
		}
	}

	return numSynced, lastBlock
}
