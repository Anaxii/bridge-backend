package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"puffinbridgebackend/wallet"
	"strconv"
)

func (h *Handler) startSync() {

	log.Info("Syncing networks")
	startBlock := map[string]int{}
	numSynced := 0
	for x := 0; numSynced < len(config.Networks); x++ {
		numSynced = 0
		for _, v := range config.Networks {
			_numSynced, lastBlock := h.sync(v, x, numSynced, h.BridgeABI)
			if x == 0 {
				startBlock[v.Name] = lastBlock
			}
			numSynced = _numSynced
		}
	}

	numSynced = 0
	for x := 0; numSynced < 1; x++ {
		_numSynced, lastBlock := h.sync(config.Subnet, x, numSynced, h.BridgeABI)
		if x == 0 {
			startBlock[config.Subnet.Name] = lastBlock
		}
		numSynced = _numSynced
	}

	for k, v := range startBlock {
		state.Write([]byte("block"), []byte(k), []byte(fmt.Sprintf("%v", v)))
	}
}

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
		nextBlock := int64(lastBlock) + 100
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
