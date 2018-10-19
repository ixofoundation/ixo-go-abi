package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
)

func main() {
	key, _ := crypto.GenerateKey()
	authorizer := bind.NewKeyedTransactor(key)
	callOpts := bind.CallOpts{
		Pending: false,
		From:    authorizer.From,
		Context: context.Background(),
	}

	url := "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk"
	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal("ERROR: %v", err)
	}

	fmt.Println("we have a connection")

	account := common.HexToAddress("0x647CD1829Ad0FF896640FCd3a29cF6Af0dE10A83")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance)

	tokenAddress := common.HexToAddress("0x827a41c26784e0f51081e6d26687bff9c1c667e6")

	ixoTokenContract, ixoError := token.NewIxoERC20Token(tokenAddress, client)

	if ixoError != nil {
		log.Fatal(ixoError)
	}

	cap, _ := ixoTokenContract.CAP(&callOpts)
	fmt.Println("IXO Token Cap:", cap)
}
