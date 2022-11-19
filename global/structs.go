package global

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type VerificationRequest struct {
	Name            string          `json:"name" bson:"name"`
	ID              string          `json:"id" bson:"id"`
	Status          string          `json:"status" bson:"status"`
	Email           string          `json:"email" bson:"email"`
	WalletAddress   string          `json:"wallet_address" bson:"wallet_address"`
	Country         string          `json:"country" bson:"country"`
	PhysicalAddress string          `json:"physical_address" bson:"physical_address"`
	IdentityNumber  string          `json:"identity_number" bson:"identity_number"`
	DateOfBirth     string          `json:"date_of_birth" bson:"date_of_birth"`
	Signature       SignatureStruct `json:"signature" bson:"signature"`
	Beneficiary     struct {
		Name          string `json:"name" bson:"name"`
		WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	}
}

type SignatureStruct struct {
	Message       string        `json:"message" bson:"message"`
	Account       string        `json:"account" bson:"account"`
	SignatureData SignatureData `json:"signature_data" bson:"signature_data"`
}

type SignatureData struct {
	HashedMessage string `json:"hashed_message" bson:"hashed_message"`
	R             string `json:"r" bson:"r"`
	S             string `json:"s" bson:"s"`
	V             string `json:"v" bson:"v"`
	Sig           string `json:"sig" bson:"sig"`
}

type ConfigStruct struct {
	PrivateKey string              `json:"private_key"`
	Networks   map[string]Networks `json:"networks"`
	Subnet     Networks            `json:"subnet"`
}

type Networks struct {
	Name          string   `json:"string"`
	RpcURL        string   `json:"rpc_url"`
	WSURL         string   `json:"ws_url"`
	ChainId       *big.Int `json:"chain_id"`
	KYCAddress    string   `json:"kyc_address"`
	BridgeAddress string   `json:"bridge_address"`
}

type NetworkLog struct {
	Log     types.Log
	Network Networks
}
