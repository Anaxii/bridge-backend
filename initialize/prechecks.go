package initialize

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"puffinbridgebackend/config"
	"time"
)

func RunPreChecks() {
	if config.PublicKey == "" {
		log.Warn("Invalid private key")
		log.Warn("Enter a new private key or edit config.json and restart:")

		var pkey string
		fmt.Scanln(&pkey)

		_publicKey, _privateKey := config.GenerateECDSAKey(pkey)
		config.PublicKey = _publicKey
		config.PrivateKey = _privateKey

		log.WithFields(log.Fields{
			"public_key": config.PublicKey,
		}).Info("New private key initialized")
	}
	log.Info("Private key check passed")

	waitForPass("balances")
	log.Info("Network balances threshold passed")

	waitForPass("kyc")
	log.Info("KYC checks passed")

	waitForPass("bridge")
	log.Info("Bridge checks passed")

	// subnet bridge permissions passed
}

func waitForPass(_type string) {
	pass := false
	for !pass {
		if _type == "balances" {
			pass = checkForBalances()
		}
		if _type == "kyc" {
			pass = checkForKYC()
		}
		if _type == "bridge" {
			pass = checkForBridgeVoter()
		}
		time.Sleep(time.Second)
	}
}

