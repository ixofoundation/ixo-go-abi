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
	ropstenAccount := generateTransactorFromPvtKey("D55B3040D0442BBBDA8593C847E2BF4561CB942C63931609BE0992B7FCB19673")
	ropstenAccount.GasLimit = uint64(5000000)

	// Create projectDid
	hexEncodedProjectDid := hex.EncodeToString([]byte("G8pj1V1Ccco7NpoGfkqY7K"))
	var projectDid [32]byte
	copy(projectDid[:], []byte("0x"+hexEncodedProjectDid))

	callOpts := bind.CallOpts{
		Pending: false,
		From:    ropstenAccount.From,
		Context: context.Background(),
	}

	url := "https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk"
	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal("ERROR: %v", err)
	}
	fmt.Println("we have a connection")

	getBalance(*client, ropstenAccount.From)
	checkTokenCap(client, callOpts)

	projectRegContract, projectError := deployProjectRegContract(client)

	if projectError != nil {
		log.Fatal("ERROR: %v", projectError)
	}

	createProjectWallet(client, ropstenAccount, projectDid, projectRegContract)
	checkProjectWallet(projectRegContract, callOpts, projectDid)
}

func generateTransactorFromPvtKey(_privateKey string) *bind.TransactOpts {
	// My Robsten Account
	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		log.Fatal(err)
	}

	return bind.NewKeyedTransactor(privateKey)
}

// Test Ropsten balance
func getBalance(client ethclient.Client, address common.Address) {
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Address balance", balance)
}

func checkTokenCap(client *ethclient.Client, callOpts bind.CallOpts) {
	tokenAddress := common.HexToAddress("0x827a41c26784e0f51081e6d26687bff9c1c667e6")
	ixoTokenContract, ixoError := token.NewIxoERC20Token(tokenAddress, client)

	if ixoError != nil {
		log.Fatal(ixoError)
	}
	cap, _ := ixoTokenContract.CAP(&callOpts)
	fmt.Println("IXO Token Cap:", cap)
}

func deployProjectRegContract(client *ethclient.Client) (*project.ProjectWalletRegistry, error) {
	projectRegAddress := common.HexToAddress("0xfe45b990a1dd890adfac13b0a9c77758cc83a862")
	return project.NewProjectWalletRegistry(projectRegAddress, client)
}

func createProjectWallet(client *ethclient.Client, ropstenAccount *bind.TransactOpts, projectDid [32]byte, projectRegContract *project.ProjectWalletRegistry) {
	fmt.Println("ProjectDid: ", string(projectDid[:]))
	transaction, _ := projectRegContract.EnsureWallet(ropstenAccount, projectDid)
	client.SendTransaction(context.Background(), transaction)
}

func checkProjectWallet(projectRegContract *project.ProjectWalletRegistry, callOpts bind.CallOpts, projectDid [32]byte) {
	walletAddress, addressErr := projectRegContract.WalletOf(&callOpts, projectDid)
	if addressErr != nil {
		log.Fatal(addressErr)
	}
	fmt.Println("Wallet address: ", walletAddress.Hex())
}
