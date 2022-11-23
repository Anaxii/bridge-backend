package events

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type LogBridgeData struct {
	User    common.Address
	Asset   common.Address
	Amount  *big.Int
	Id      [32]byte
	ChainId *big.Int
	Block   int64
}

type LogBridgeCancelData struct {
	User  common.Address
	Id    [32]byte
	Block int64
}
