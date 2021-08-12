# Take Home Exercise

`cocsethe` stands for `Crypto.Org Chain SDK Engineer Take Home Exercise`

In this project, I implement a simple wrapper library to perform transaction generation on Ethereum network(it supports the EIP 1559 transaction type).

Others can use this library(`cocsethe`) to generate transaction and query balance.


## Usage

```go
// generate a new Client with given rpcURL
client, _ := NewEtherClient(rpcURL)

// construct transaction options
options := TransactionOptions{
    FromAddressHexKey: privateKey,
    ToAddress:         toAddress,
    Amount:            100000000000000, // 0.0001
}

// make a transaction
transactionHash, err := client.GenerateTransaction(options)

if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("transaction hash: %s", *transactionHash)

// query transaction
transaction, isPending, err := client.GetTransaction("0x020d79082faa2e8f44a5f3e17a834775c54d9b96d78c460e6e695f1eb3a0faca")
if err != nil {
	fmt.Println(err)
    return
}

fmt.Printf("transaction: %s, isPending: %t", transaction.Hash(), isPending)
```

## 3rd party libraries used

- [go-ethereum](https://github.com/ethereum/go-ethereum) to interact with the ethereum network
  - It's not only ethereum client, but also a go library

## Security

To generate a transaction, user have to pass in the private key,
but in this library, we don't store it or send it via network,
just use the key to sign the transaction, so it is safe to use.

## Coding

- Use factory pattern to generate `Client` & `Account`
- Use interface to abstract `Client` and implement it with an `etherClient` entry
- Write testing to ensure the function is working and user can figure how to use it

## Testing

I'm using Rinkeby Test Network for testing.

- RPC URL: https://rinkeby.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161
- etherscan: https://rinkeby.etherscan.io/

I also created a test account,

- address: `0x411e2725F2Cd5358AEe921bEB4De8299BC0181EB`
- private key: `08346faf1f20d2dcbf63887ca87ffcc211809cad438fbe6e62c587403d2bf4fa`