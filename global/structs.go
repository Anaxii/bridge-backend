package global

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type ConfigStruct struct {
	PrivateKey string              `json:"private_key"`
	APIPort    string              `json:"api_port"`
	APIEnabled bool                `json:"api_enabled"`
	Networks   map[string]Networks `json:"networks"`
	Subnet     Networks            `json:"subnet"`
}

type Networks struct {
	Name             string   `json:"string"`
	RpcURL           string   `json:"rpc_url"`
	WSURL            string   `json:"ws_url"`
	ChainId          *big.Int `json:"chain_id"`
	KYCAddress       string   `json:"kyc_address"`
	BridgeAddress    string   `json:"bridge_address"`
	BlockRequirement int      `json:"block_requirement"`
}

type NetworkLog struct {
	Log     types.Log
	Network Networks
}
