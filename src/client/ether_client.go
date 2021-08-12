package client

import (
	"COCSETHE/src/model"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// NewEtherClient creates a Client for give rpc url
func NewEtherClient(url string) (Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	return etherClient{client: client}, nil
}

// etherClient implements Client interface
type etherClient struct {
	client *ethclient.Client
}

// QueryBalance query balance for the give address
func (e etherClient) QueryBalance(address string) (*model.Balance, error) {
	account := common.HexToAddress(address)
	balance, err := e.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	return &model.Balance{
		Amount: balance,
		Unit:   "wei",
	}, nil
}

// GenerateTransaction generate transaction for give options, transaction is created and signed on the fly
// it returns generated transaction hash, you can take it for further use, eg, GetTransaction
func (e etherClient) GenerateTransaction(options TransactionOptions) (*string, error) {
	chainID, err := e.getChainID()
	if err != nil {
		return nil, err
	}

	account, err := model.NewEtherAccount(options.FromAddressHexKey)
	if err != nil {
		return nil, err
	}

	nonce, err := e.client.PendingNonceAt(context.Background(), common.Address(account.Address))
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(options.Amount) // in wei

	gas := uint64(21000) // in units

	toAddress := common.HexToAddress(options.ToAddress)

	gasTipCap, _ := e.client.SuggestGasTipCap(context.Background())

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: gasTipCap,
		GasTipCap: gasTipCap,
		Gas:       gas,
		To:        &toAddress,
		Value:     value,
		Data:      nil,
	})

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), account.PrivateKey)
	if err != nil {
		return nil, err
	}

	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	transactionID := signedTx.Hash().Hex()

	return &transactionID, nil
}

// GetTransaction query transaction for the given transactionHash
func (e etherClient) GetTransaction(transactionHash string) (*types.Transaction, bool, error) {
	hash := common.HexToHash(transactionHash)
	return e.client.TransactionByHash(context.Background(), hash)
}

// getChainID is a helper method for client to get chainID
func (e etherClient) getChainID() (*big.Int, error) {
	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	return chainID, nil
}
