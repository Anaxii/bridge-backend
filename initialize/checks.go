package initialize

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/big"
	"puffinbridgebackend/blockchain/contractInteraction"
	"puffinbridgebackend/config"
	"puffinbridgebackend/wallet"
)

func checkForBalances() bool {
	for _, v := range config.Networks {
		bal, _ := wallet.Balance(v)
		if bal.Cmp(big.NewInt(5)) < 0 {
			log.WithFields(log.Fields{
				"balance":        bal.String(),
				"threshold":      fmt.Sprintf("%v", 5),
				"network":        v.Name,
				"wallet_address": config.PublicKey,
			}).Warn("Balance too low, waiting for more funds.")
			return false
		}
	}
	bal, _ := wallet.Balance(config.Subnet)
	if bal.Cmp(big.NewInt(5)) < 0 {
		log.WithFields(log.Fields{
			"balance":        bal.String(),
			"threshold":      fmt.Sprintf("%v", 5),
			"network":        "subnet",
			"wallet_address": config.PublicKey,
		}).Warn("Balance too low, waiting for more funds.")
		return false
	}
	return true
}

func checkForKYC() bool {
	kyc, msg := hasKYC()
	if !kyc {
		log.WithFields(log.Fields{
			"network":        msg,
			"wallet_address": config.PublicKey,
		}).Warn("Wallet not approved on network.")
		return false
	}
	return true
}

func checkForBridgeVoter() bool {
	kyc, msg := hasBridgeVotingPower()
	if !kyc {
		log.WithFields(log.Fields{
			"network":        msg,
			"wallet_address": config.PublicKey,
		}).Warn("Wallet not voter on network bridge.")
		return false
	}
	return true
}

func hasKYC() (bool, string) {
	for _, v := range config.Networks {
		if !contractInteraction.IsKYCOnMainnet(v) {
			log.WithFields(log.Fields{
				"wallet_address": config.PublicKey,
				"network":        v.Name,
			}).Warn("Wallet does not have kyc on mainnet.")
			return false, v.Name
		}
	}
	if !contractInteraction.IsKYCOnSubnet(config.Subnet) {
		log.WithFields(log.Fields{
			"wallet_address": config.PublicKey,
		}).Warn("Wallet does not have kyc on subnet.")
		return false, config.Subnet.Name
	}
	return true, ""
}

func hasBridgeVotingPower() (bool, string) {
	for _, v := range config.Networks {
		if !contractInteraction.IsVoterOnMainnet(v) {
			log.WithFields(log.Fields{
				"wallet_address": config.PublicKey,
				"network":        v.Name,
			}).Warn("Wallet does not have voting power on mainnet.")
			return false, v.Name
		}
	}
	if !contractInteraction.IsVoterOnMainnet(config.Subnet) {
		log.WithFields(log.Fields{
			"wallet_address": config.PublicKey,
		}).Warn("Wallet does not have voting power on subnet.")
		return false, config.Subnet.Name
	}
	return true, ""
}
