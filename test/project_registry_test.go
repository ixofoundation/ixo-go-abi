package test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	auth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
)

func TestProjectRegistryContract(t *testing.T) {
	key, _ := crypto.GenerateKey()
	authorizer := bind.NewKeyedTransactor(key)
	blockchain := backends.NewSimulatedBackend(core.GenesisAlloc{authorizer.From: {Balance: big.NewInt(10000000000)}}, uint64(10000000))

	hexEncodedProjectDid := hex.EncodeToString([]byte("G8pj1V1Bcco7NpoGfkqY7K"))
	t.Logf("ProjectDid: %v", hexEncodedProjectDid)

	var projectDid [32]byte
	copy(projectDid[:], []byte("0x" + hexEncodedProjectDid))
	t.Logf("ProjectDid: %v", string(projectDid[:]))

	callOpts := bind.CallOpts{
		Pending: true,
		From:    authorizer.From,
	}

	transOpts := bind.TransactOpts{
		From:     authorizer.From,
		Signer:   authorizer.Signer,
		GasLimit: uint64(100000),
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

	blockchain.Commit()

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

	blockchain.Commit()

	t.Run("Renounce ownership", renounceOwnership(*projectRegistryContact, transOpts))
	blockchain.Commit()
	t.Run("Check ownership", checkOwnership(*projectRegistryContact, callOpts))
	t.Run("Create project wallet", createProjectWallet(*projectRegistryContact, transOpts, projectDid))
	blockchain.Commit()
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

		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)

		t.Logf("Wallet address: %v", projectWalletAddress.String())

		// TODO: Fix
		if projectWalletAddress.String() == "0x0000000000000000000000000000000000000000" {
			t.Logf("Project Wallet: %v", projectWalletAddress.String())
			t.Errorf("Expected owner to be: 0x0000000000000000000000000000000000000000")
		}

	}
}
