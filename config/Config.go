package config

import (
	"crypto/ecdsa"
	"encoding/json"
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

func init() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		file, _ := json.MarshalIndent(global.ConfigStruct{
			Networks: map[string]global.Networks{
				"fuji": {
					Name:          "fuji",
					RpcURL:        "https://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc",
					WSURL:        "wss://red-weathered-firefly.avalanche-testnet.quiknode.pro/ext/bc/C/rpc",
					ChainId:       big.NewInt(43113),
					KYCAddress:    "0x094B85f01716ddB7E07bE8E68c29d1bA6E59944e",
					BridgeAddress: "0x7D761A316dbD5d6Ae69cb7B50b234117716a26b8",
								},
			},
			Subnet:   global.Networks{
				Name: "puffin",
				RpcURL: "https://node.thepuffin.network/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/rpc",
				WSURL: "wss://node.thepuffin.network/ext/bc/273dwzFtrR6JQzLncTAbN5RBtiqdysVfKTJKBvYHhtUHBnrYWe/rpc",
				ChainId: big.NewInt(43113114),
				KYCAddress: "0x0200000000000000000000000000000000000002",
			},
		}, "", "  ")
		_ = ioutil.WriteFile("config.json", file, 0644)
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

	Networks = config.Networks
	Subnet = config.Subnet
	if config.PrivateKey != "" {
		_publicKey, _privateKey := GenerateECDSAKey(config.PrivateKey)
		PublicKey = _publicKey
		PrivateKey = _privateKey
	}

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
