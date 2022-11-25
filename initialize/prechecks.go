package initialize

import (
	log "github.com/sirupsen/logrus"
	"os"
	"puffinbridgebackend/config"
	"time"
)

func RunPreChecks() {
	if config.PublicKey == "" {

		log.Warn("No private key found in config.json")

		var pkey string

		key := os.Getenv("private_key")
		if key != "" {
			pkey = key
			log.WithFields(log.Fields{
				"public_key": config.PublicKey,
			}).Info("Using private key from env")
		} else {
			log.Fatal("Edit config.json or set env variable private_key")
		}

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
