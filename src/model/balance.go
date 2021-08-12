package model

import (
	"github.com/ethereum/go-ethereum/params"
	"math/big"
)

// Balance is the QueryBalance result
type Balance struct {
	Amount *big.Int // Amount of the Balance
	Unit   string   // Unit in wei
}

// ToEther converts Amount of Balance in wei to ether
func (b Balance) ToEther() *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(b.Amount), big.NewFloat(params.Ether))
}
