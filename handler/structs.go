package handler

import (
	ethABI "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"puffinbridgebackend/global"
)

type Handler struct {
	BridgeQueue []BridgeRequest
	BridgeABI   ethABI.ABI
	Blocks      map[string]int
	Logs        []LogHistory
}

type LogHistory struct {
	Log       BridgeRequest `json:"log"`
	Error     string        `json:"error"`
	Status    string        `json:"status"`
	Timestamp int64         `json:"timestamp"`
}

type BridgeRequest struct {
	Id         [32]byte
	NetworkIn  global.Networks
	NetworkOut global.Networks
	User       common.Address
	Asset      common.Address
	Amount     *big.Int
	Block      int64
	Method     string
}
