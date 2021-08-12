package client

import (
	"fmt"
	"testing"
)

var (
	rpcURL     = "https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
	privateKey = "08346faf1f20d2dcbf63887ca87ffcc211809cad438fbe6e62c587403d2bf4fa"
	toAddress  = "0xBAe9d2c20D602509C62e9F34e562791d8e8Ec53f"
)

func TestCreateANewEtherClientWillNotFail(t *testing.T) {
	_, err := NewEtherClient(rpcURL)

	if err != nil {
		t.Fatal("Create client failed")
	}
}

func TestBalanceToReadableEtherWillNotFail(t *testing.T) {
	client, _ := NewEtherClient(rpcURL)

	balance, err := client.QueryBalance("0xBAe9d2c20D602509C62e9F34e562791d8e8Ec53f")

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("balance: %d %s\n", balance.Amount, balance.Unit)
}

func TestGenerateTransactionWillNotFail(t *testing.T) {
	client, _ := NewEtherClient(rpcURL)

	options := TransactionOptions{
		FromAddressHexKey: privateKey,
		ToAddress:         toAddress,
		Amount:            100000000000000, // 0.0001
	}

	transactionHash, err := client.GenerateTransaction(options)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("transaction hash: %s", *transactionHash)
}

func TestGetTransactionNotFail(t *testing.T) {
	client, _ := NewEtherClient(rpcURL)

	transaction, isPending, err := client.GetTransaction("0x020d79082faa2e8f44a5f3e17a834775c54d9b96d78c460e6e695f1eb3a0faca")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("transaction: %s, isPending: %t", transaction.Hash(), isPending)

}
