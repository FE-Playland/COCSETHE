package model

import (
	gocrypto "crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// AddressLength is the expected length of the address
	AddressLength = 20
)

// Address is the ether address in bytes
type Address [AddressLength]byte

// Account is the ether account, created with private key in hex format, Address will be generated from PrivateKey
type Account struct {
	PrivateKey *gocrypto.PrivateKey
	Address    Address
}

// NewEtherAccount creates an Account with private key in hex format
func NewEtherAccount(hexKey string) (*Account, error) {

	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*gocrypto.PublicKey)

	if !ok {
		return nil, err
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Account{
		PrivateKey: privateKey,
		Address:    (Address)(address),
	}, nil
}
