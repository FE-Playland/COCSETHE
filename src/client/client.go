package client

import (
	"COCSETHE/src/model"
	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionOptions is used to make a transaction
type TransactionOptions struct {
	FromAddressHexKey string // transfer account private key in hex format
	ToAddress         string // address to receive transfer
	Amount            int64  // unit is wei
}

// Client is the main entry for user to interact with the ethereum, it's an abstract interface
type Client interface {
	// QueryBalance query balance for the give address
	QueryBalance(address string) (*model.Balance, error)
	// GenerateTransaction generate transaction for give options, transaction is created and signed on the fly
	// it returns generated transaction hash, you can take it for further use, eg, GetTransaction
	GenerateTransaction(options TransactionOptions) (*string, error)
	// GetTransaction query transaction for the given transactionHash
	GetTransaction(transactionHash string) (*types.Transaction, bool, error)
}
