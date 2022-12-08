package global

import "github.com/ethereum/go-ethereum/crypto"

var logBridgeInSig = []byte("BridgeIn(address,address,uint256,bytes32,uint256)")
var logBridgeOutSig = []byte("BridgeOut(address,address,uint256,bytes32,uint256)")
var logBridgeOutWarmSig = []byte("BridgeOutWarm(address,address,uint256,bytes32,uint256)")
var logBridgeOutCanceledSig = []byte("BridgeOutCanceled(address,bytes32)")

var LogBridgeInSigHash = crypto.Keccak256Hash(logBridgeInSig)
var LogBridgeOutSigHash = crypto.Keccak256Hash(logBridgeOutSig)
var LogBridgeOutWarmSigHash = crypto.Keccak256Hash(logBridgeOutWarmSig)
var LogBridgeOutCanceledSigHash = crypto.Keccak256Hash(logBridgeOutCanceledSig)

var Logs []interface{}
var SocketChannels =  make(map[string]chan interface{})

func Log(data interface{}) {
	for  k := range SocketChannels {
		SocketChannels[k] <- data
	}
}
