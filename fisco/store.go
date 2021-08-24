package fisco

import (
	"backend_ft_tcs/fisco/contracts/store"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func (f fiscoService) DeployStore() (string, string) {
	input := "Store deployment 1.0"
	address, tx, _, err := store.DeployStore(f.client.GetTransactOpts(), f.client, input)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved, will use in next example
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	return address.Hex(), tx.Hash().Hex()
}
func (f fiscoService) StoreSetItem(address string) {
	contractAddress := common.HexToAddress(address)
	instance, err := store.NewStore(contractAddress, f.client)
	if err != nil {
		logger.Fatal(err)
	}
	storeSession := &store.StoreSession{Contract: instance, CallOpts: *f.client.GetCallOpts(), TransactOpts: *f.client.GetTransactOpts()}

	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("kwinin"))
	tx, receipt, err := storeSession.SetItem(key, value)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
}

func (f fiscoService) StoreItems(address string) {
	contractAddress := common.HexToAddress(address)
	instance, err := store.NewStore(contractAddress, f.client)
	if err != nil {
		logger.Fatal(err)
	}
	storeSession := &store.StoreSession{Contract: instance, CallOpts: *f.client.GetCallOpts(), TransactOpts: *f.client.GetTransactOpts()}
	key := [32]byte{}
	copy(key[:], []byte("foo"))
	// read the result
	result, err := storeSession.Items(key)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Println("get item: " + string(result[:])) // "bar"
}
