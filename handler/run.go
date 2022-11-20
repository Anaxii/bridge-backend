package handler

import (
	"fmt"
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/contractABI"
	"puffinbridgebackend/events"
	"puffinbridgebackend/global"
	"puffinbridgebackend/state"
	"puffinbridgebackend/wallet"
	"strconv"
	"strings"
)

func RunHandler() {
	listenToEvents()
}

func listenToEvents() {

	var event = make(chan global.NetworkLog)

	bridgeABI, err := ethABI.JSON(strings.NewReader(abi.PuffinMainnetBridgeABI))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Syncing networks")
	numSynced := 0
	x := 0
	for numSynced < len(config.Networks) {
		numSynced = 0
		x++
		for _, v := range config.Networks {
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
			} else {
				nextBlock := int64(lastBlock) + 10
				if nextBlock > walletBlock.Int64() {
					nextBlock = walletBlock.Int64()
				}
				contractInteraction.QueryEvent(v, int64(lastBlock), nextBlock, v.BridgeAddress, bridgeABI)
				err = state.Write([]byte("block"), []byte(v.Name), []byte(fmt.Sprintf("%v", nextBlock)))
				if err != nil {
					log.WithFields(log.Fields{
						"status": "syncing",
						"block":  string(block),
					}).Warn("Could not update database")
				}
			}
		}
	}

	log.Info("Starting listeners")
	for _, v := range config.Networks {
		go contractInteraction.ListenToEvents(v, v.BridgeAddress, event)
	}

	for {
		select {
		case vLog := <-event:
			switch vLog.Log.Topics[0].Hex() {
			case global.LogBridgeInSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeIn", vLog.Log)
			case global.LogBridgeOutSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeOut", vLog.Log)
			case global.LogBridgeOutWarmSigHash.Hex():
				_, err = events.LogBridge(bridgeABI, "BridgeOutWarm", vLog.Log)
			case global.LogBridgeOutCanceledSigHash.Hex():
				_, err = events.LogBridgeCancel(bridgeABI, "BridgeCancel", vLog.Log)
			}
			state.Write([]byte("block"), []byte(vLog.Network.Name), []byte(fmt.Sprintf("%v", vLog.Log.BlockNumber)))

		}
	}
}
