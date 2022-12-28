package initialize

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func RunPreChecks() {

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
