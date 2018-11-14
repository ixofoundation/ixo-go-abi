package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	auth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
)

func TestProjectRegistryContract(t *testing.T) {
	// Use Ganache Account
	key, _ := crypto.HexToECDSA("b6ad4d7b59a2766e94f9290740fd62676165684500c6d1331185912600e19481")
	authorizer := bind.NewKeyedTransactor(key)

	// Use Ganache 
	blockchain, _ := ethclient.Dial("http://127.0.0.1:7545")

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
		From:     authorizer.From,
		Signer:   authorizer.Signer,
	}

	// Deploy token contract
	tokenContractAddress, _, _, _ := project.DeployIxoERC20Token(
		authorizer,
		blockchain,
	)

	// Setup quorum
	var quorum []common.Address
	quorum = append(quorum, authorizer.From)

	// Deploy auth contract
	authContractAddress, _, _, _ := auth.DeployAuthContract(
		authorizer,
		blockchain,
		quorum,
		big.NewInt(1),
	)

	// Deploy factory contract
	factoryContractAddress, _, _, _ := project.DeployProjectWalletFactory(
		authorizer,
		blockchain,
	)

	// Deploy project registry contract
	_, _, projectRegistryContact, regDeployErr := project.DeployProjectWalletRegistry(
		authorizer,
		blockchain,
		tokenContractAddress,
		authContractAddress,
		factoryContractAddress,
	)
	if regDeployErr != nil {
		t.Errorf("ERROR: %v", regDeployErr)
	}

	t.Run("Renounce ownership", renounceOwnership(*projectRegistryContact, transOpts))
	t.Run("Check ownership", checkOwnership(*projectRegistryContact, callOpts))
	t.Run("Create project wallet", createProjectWallet(*projectRegistryContact, transOpts, projectDid))
	t.Run("Check if project wallet exist", checkProjectWallet(*projectRegistryContact, callOpts, projectDid))
}

// util function to log transaction response
func logTxnResponse(t *testing.T, transaction types.Transaction) {
	txnResponse, _ := transaction.MarshalJSON()
	response := string(txnResponse[:])
	t.Logf("Transaction Response: %v", response)
}

func renounceOwnership(projectRegistryContact project.ProjectWalletRegistry, transOpts bind.TransactOpts) func(*testing.T) {
	return func(t *testing.T) {
		transaction, renounceErr := projectRegistryContact.RenounceOwnership(&transOpts)

		if renounceErr != nil {
			t.Errorf("ERROR: %v", renounceErr)
		}

		logTxnResponse(t, *transaction)
	}
}

func checkOwnership(projectRegistryContact project.ProjectWalletRegistry, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {

		if owner, _ := projectRegistryContact.Owner(&callOpts); owner.String() != "0x0000000000000000000000000000000000000000" {
			t.Errorf("Expected owner to be: 0x0000000000000000000000000000000000000000")
		}

	}
}

func createProjectWallet(projectRegistryContact project.ProjectWalletRegistry, transOpts bind.TransactOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {

		t.Logf("Ensure Wallet: %v", string(projectDid[:]))

		// create project wallet
		transaction, walletErr := projectRegistryContact.EnsureWallet(&transOpts, projectDid)

		if walletErr != nil {
			t.Errorf("ERROR: %v", walletErr)
		}
		logTxnResponse(t, *transaction)
	}
}

func checkProjectWallet(projectRegistryContact project.ProjectWalletRegistry, callOpts bind.CallOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {

		t.Logf("WalletOf: %v", string(projectDid[:]))

		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)

		t.Logf("Wallet address: %v", projectWalletAddress.String())

		// TODO: Fix
		if projectWalletAddress.String() == "0x0000000000000000000000000000000000000000" {
			t.Logf("Project Wallet: %v", projectWalletAddress.String())
			t.Errorf("Expected owner to be: 0x0000000000000000000000000000000000000000")
		}

	}
}
