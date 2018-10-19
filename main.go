package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
)

func main() {

	// My Robsten Account
	privateKey, err := crypto.HexToECDSA("D55B3040D0442BBBDA8593C847E2BF4561CB942C63931609BE0992B7FCB19673")
	if err != nil {
		log.Fatal(err)
	}

	myRopstenAccount := bind.NewKeyedTransactor(privateKey)

	callOpts := bind.CallOpts{
		Pending: false,
		From:    myRopstenAccount.From,
		Context: context.Background(),
	}

	transOpts := bind.TransactOpts{
		From:     myRopstenAccount.From,
		Signer:   myRopstenAccount.Signer,
		GasLimit: uint64(5000000),
	}

	url := "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk"
	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal("ERROR: %v", err)
	}

	fmt.Println("we have a connection")

	// Test Ropsten balance
	account := common.HexToAddress("0x647CD1829Ad0FF896640FCd3a29cF6Af0dE10A83")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Address balance", balance)

	// ERC20
	tokenAddress := common.HexToAddress("0x827a41c26784e0f51081e6d26687bff9c1c667e6")
	ixoTokenContract, ixoError := token.NewIxoERC20Token(tokenAddress, client)

	if ixoError != nil {
		log.Fatal(ixoError)
	}
	cap, _ := ixoTokenContract.CAP(&callOpts)
	fmt.Println("IXO Token Cap:", cap)

	// Project
	hexEncodedProjectDid := hex.EncodeToString([]byte("G8pj1V1Bcco7NpoGfkqY7K"))
	var projectDid [32]byte
	copy(projectDid[:], []byte("0x"+hexEncodedProjectDid))

	projectRegAddress := common.HexToAddress("0xfe45b990a1dd890adfac13b0a9c77758cc83a862")
	projectRegContract, projectError := project.NewProjectWalletRegistry(projectRegAddress, client)

	if projectError != nil {
		log.Fatal(projectError)
	}

	fmt.Println("ProjectDid: ", string(projectDid[:]));

	transaction, _ := projectRegContract.EnsureWallet(&transOpts, projectDid)

	client.SendTransaction(context.Background(), transaction)

	walletAddress, _ := projectRegContract.WalletOf(&callOpts, projectDid)

	fmt.Println("Wallet address: ", walletAddress.Hex());
}
