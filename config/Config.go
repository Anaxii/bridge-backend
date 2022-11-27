package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"math/big"
	"os"
	"puffinbridgebackend/global"
)

var PrivateKey *ecdsa.PrivateKey
var PublicKey string
var Networks map[string]global.Networks
var Subnet global.Networks
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

		file, _ := json.MarshalIndent(global.ConfigStruct{
			PrivateKey: fmt.Sprintf("%v", hexutil.Encode(privateKeyBytes)[2:]),
			APIPort:    "8080",
			APIEnabled: true,
			Networks: map[string]global.Networks{
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
			Subnet: global.Networks{
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

	var config global.ConfigStruct
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("Could not parse config", err)
	}

	key := os.Getenv("private_key")
	if key != "" {
		config.PrivateKey = key
	}

	Networks = config.Networks
	Subnet = config.Subnet
	if config.PrivateKey != "" {
		_publicKey, _privateKey := GenerateECDSAKey(config.PrivateKey)
		PublicKey = _publicKey
		PrivateKey = _privateKey
	}
	APIPort = config.APIPort
	APIEnabled = config.APIEnabled

	log.Info("Config initialized")

}

func GenerateECDSAKey(pkey string) (string, *ecdsa.PrivateKey) {
	privateKey, err := crypto.HexToECDSA(pkey)
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	_publicKey := hexutil.Encode(hash.Sum(nil)[12:])
	_privateKey := privateKey

	return _publicKey, _privateKey
}
