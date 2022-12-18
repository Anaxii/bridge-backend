package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"os"
	"puffinbridgebackend/util"
)

var PrivateKey *ecdsa.PrivateKey
var PublicKey string
var NetworksMap map[string]Networks
var Subnet Networks
var APIPort string
var APIEnabled bool

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {

		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)

		fmt.Println("Enter private key passphrase: ")
		var k string
		fmt.Scanln(&k)

		util.EncryptAES([]byte(k), []byte(hexutil.Encode(privateKeyBytes)[2:]))

		file, _ := json.MarshalIndent(ConfigStruct{
			APIPort:    "1234",
			APIEnabled: true,
			Networks: map[string]Networks{
				"fuji": {
					Name:             "fuji",
					RpcURL:           "https://node.thepuffin.network/ext/bc/C/rpc",
					WSURL:            "ws://52.35.42.217:9650/ext/bc/C/ws",
					ChainId:          big.NewInt(43113),
					KYCAddress:       "0x0d3378f28cccF59E81084Cf4f676cBcaB64ca359",
					BridgeAddress:    "0x40dA58598877009868B9B876f52d31a0C204FC63",
					BlockRequirement: 3,
				},
			},
			Subnet: Networks{
				Name:             "puffin",
				RpcURL:           "https://node.thepuffin.network/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/rpc",
				WSURL:            "ws://52.35.42.217:9650/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/ws",
				ChainId:          big.NewInt(43113114),
				KYCAddress:       "0x0200000000000000000000000000000000000002",
				BridgeAddress:    "0xd3E11DeF6d34E231ab410e5aA187e1f2d9fF19E1",
				BlockRequirement: 0,
			},
		}, "", "  ")
		_ = ioutil.WriteFile("config.json", file, 0644)
		jsonFile, err = os.Open("config.json")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Config JSON invalid", err)
	}

	var config ConfigStruct
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("Could not parse config", err)
	}

	key := os.Getenv("private_key")
	if key != "" {
		config.PrivateKey = key
	}

	NetworksMap = config.Networks
	Subnet = config.Subnet

	_publicKey, _privateKey := util.GenerateECDSAKey(util.DecryptAES([]byte(os.Getenv("pass"))))
	PublicKey = _publicKey
	PrivateKey = _privateKey
	APIPort = config.APIPort
	APIEnabled = config.APIEnabled

	log.Info("Config initialized")

}


