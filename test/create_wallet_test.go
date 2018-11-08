package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
)

func TestProjectRegistry(t *testing.T) {
	// Use Ganache Account
	key, _ := crypto.HexToECDSA("F16653CA1D9E207CA68DE145C89D89F3E899B4BADEAC726EA9F62AEA0AB1CAA9")
	authorizer := bind.NewKeyedTransactor(key)

	// Use Ganache
	blockchain, _ := ethclient.Dial("https://ropsten.infura.io/sq19XM5Eu2ANGAzwZ4yk")

	did := "WjU6gE1JhZANcdv3aC8PEJ"

	var projectDid [32]byte
	copy(projectDid[:], did)
	t.Logf("ProjectDid Byte Array: %v", projectDid)
	t.Logf("ProjectDid Byte Array: %v", string(projectDid[:]))

	callOpts := bind.CallOpts{
		Pending: true,
		From:    authorizer.From,
	}

	transOpts := bind.TransactOpts{
		From:   authorizer.From,
		Signer: authorizer.Signer,
	}

	// Deploy project registry contract
	projectRegistryContact, regDeployErr := project.NewProjectWalletRegistry(
		common.HexToAddress("0xa68b5c41f601594763e09cf1d51943d96e34f81d"),
		blockchain,
	)
	if regDeployErr != nil {
		t.Errorf("ERROR: %v", regDeployErr)
	}

	//blockchain.Commit()

	t.Run("Create Project Wallet", ensureWallet(*projectRegistryContact, transOpts, projectDid))
	t.Run("Check if project wallet exist", getProjectWallet(*projectRegistryContact, callOpts, projectDid))
}

// util function to log transaction response
func logTxnResp(t *testing.T, transaction types.Transaction) {
	txnResponse, _ := transaction.MarshalJSON()
	response := string(txnResponse[:])
	t.Logf("Transaction Response: %v", response)
}

func ensureWallet(projectRegistryContact project.ProjectWalletRegistry, transOpts bind.TransactOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {
		transaction, renounceErr := projectRegistryContact.EnsureWallet(&transOpts, projectDid)

		if renounceErr != nil {
			t.Errorf("ERROR: %v", renounceErr)
		}

		logTxnResp(t, *transaction)
	}
}

func getProjectWallet(projectRegistryContact project.ProjectWalletRegistry, callOpts bind.CallOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {

		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)

		t.Logf("Wallet address: %v", projectWalletAddress.String())

		// TODO: Fix
		if projectWalletAddress.String() == "0x0000000000000000000000000000000000000000" {
			t.Logf("Project Wallet: %v", projectWalletAddress.String())
			t.Errorf("Expected owner to be: 0x0000000000000000000000000000000000000000")
		}

	}
}
