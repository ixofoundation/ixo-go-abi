package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	auth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
	util "github.com/ixofoundation/ixo-go-abi/test/util"
	"github.com/stretchr/testify/assert"
)

func TestProjectRegistryContract(t *testing.T) {
	// Use Ganache Account
	// Contract Owner Address
	keyAuth, _ := crypto.HexToECDSA("d4291707c68a3e43e7f9a605f46d8bc2730af5023046a905b7028469d49cb868")
	ownerWallet := bind.NewKeyedTransactor(keyAuth)
	ownerWallet.GasLimit = uint64(27821000)

	// Evaluator Address
	keyEval, _ := crypto.HexToECDSA("baf7af00a5f868db6ef8ca22ebbf69c131217ef08427c040d90931a803d98957")
	evaluatorWallet := bind.NewKeyedTransactor(keyEval)

	// Validator Address
	keyValidator, _ := crypto.HexToECDSA("de9e861326d46d132312bc140468614a2d6b0c41ad6801f933d77129a1be4d4e")
	validatorWallet := bind.NewKeyedTransactor(keyValidator)
	validatorWallet.GasLimit = uint64(27821000) // VERY IMPORTANT!!!!!

	// Use Ganache
	blockchain, _ := ethclient.Dial("http://127.0.0.1:7545")

	var projectDid [32]byte
	copy(projectDid[:], "WjU6gE1JhZANcdv3aC8PEJ")

	callOpts := bind.CallOpts{
		Pending: true,
		From:    ownerWallet.From,
	}

	// Deploy token contract
	ixoTokenAddress, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		ownerWallet,
		blockchain,
	)
	t.Logf("ERC20_ADDRESS: %v", ixoTokenAddress.Hex())

	// Set minter
	ixoTokenContact.SetMinter(ownerWallet, ownerWallet.From)

	// Setup quorum
	members := []common.Address{validatorWallet.From}

	// Deploy auth contract
	authAddress, _, authContract, _ := auth.DeployAuthContract(
		ownerWallet,
		blockchain,
		members,
		big.NewInt(1),
	)
	t.Logf("AUTH_CONTRACT_ADDRESS: %v", authAddress.Hex())

	// Deploy factory contract
	factoryContractAddress, _, _, _ := project.DeployProjectWalletFactory(
		ownerWallet,
		blockchain,
	)

	// Deploy project registry contract
	_, _, projectRegistryContact, _ := project.DeployProjectWalletRegistry(
		ownerWallet,
		blockchain,
		ixoTokenAddress,
		authAddress,
		factoryContractAddress,
	)

	projectRegistryContact.SetFactory(ownerWallet, factoryContractAddress)

	// DEPLOY_PROJECT_WALLET_AUTH_CONTRACT
	projectWalletAuthAddress, _, projectWalletAuthContract, _ := auth.DeployProjectWalletAuthoriser(ownerWallet, blockchain)
	t.Logf("PROJECT_WALLET_AUTH_ADDRESS: %v", projectWalletAuthAddress.Hex())

	// SET_PROJECT_WALLET_AUTH_OWNER
	projectWalletAuthContract.SetAuthoriser(ownerWallet, authAddress)

	t.Run("Renounce ownership", renounceOwnership(*projectRegistryContact, *ownerWallet, callOpts))
	t.Run("Create project wallet", createProjectWallet(*projectRegistryContact, *ownerWallet, callOpts, projectDid))
	t.Run("Fund project wallet", fundProjectWallet(*projectRegistryContact, *ixoTokenContact, *ownerWallet, callOpts, projectDid))
	t.Run("Pay evaluator from auth contract", payEvaluatorFromAuthContract(*projectRegistryContact, *authContract, projectWalletAuthAddress, evaluatorWallet.From, *validatorWallet, callOpts, projectDid, util.Random32Bytes()))
}

func renounceOwnership(projectRegistryContact project.ProjectWalletRegistry, transOpts bind.TransactOpts, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		projectRegistryContact.RenounceOwnership(&transOpts)
		owner, _ := projectRegistryContact.Owner(&callOpts)
		assert.EqualValues(t, owner.Hex(), "0x0000000000000000000000000000000000000000", "Owner should be: 0x0000000000000000000000000000000000000000")
	}
}

func createProjectWallet(projectRegistryContact project.ProjectWalletRegistry, transOpts bind.TransactOpts, callOpts bind.CallOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {
		projectRegistryContact.EnsureWallet(&transOpts, projectDid)
		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)
		assert.NotEqual(t, projectWalletAddress, "0x0000000000000000000000000000000000000000", "Error creating project wallet!")
	}
}

func fundProjectWallet(projectRegistryContact project.ProjectWalletRegistry, ixoTokenContract token.IxoERC20Token, ownerWallet bind.TransactOpts, callOpts bind.CallOpts, projectDid [32]byte) func(*testing.T) {
	return func(t *testing.T) {
		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)
		t.Logf("walletAddress: %v", projectWalletAddress.Hex())
		ixoTokenContract.Mint(&ownerWallet, projectWalletAddress, big.NewInt(500000000))
		projectWalletBalance, _ := ixoTokenContract.BalanceOf(&callOpts, projectWalletAddress)
		assert.EqualValues(t, big.NewInt(500000000), projectWalletBalance, "Error while funding project wallet!")
	}
}

func payEvaluatorFromAuthContract(projectRegistryContact project.ProjectWalletRegistry, authContract auth.AuthContract, projectWalletAuthAddress common.Address, evaluatorAddress common.Address, validatorWallet bind.TransactOpts, callOpts bind.CallOpts, projectDid [32]byte, txID [32]byte) func(*testing.T) {
	return func(t *testing.T) {

		projectWalletAddress, _ := projectRegistryContact.WalletOf(&callOpts, projectDid)

		t.Logf("WALLET_ADDRESS: %v", projectWalletAddress.Hex())
		t.Logf("VALIDATOR_NODE_WALLET: %v", validatorWallet.From.Hex())
		t.Logf("PROJECT_WALLET_AUTH: %v", projectWalletAuthAddress.Hex())
		t.Logf("EVALUATOR_WALLET: %v", evaluatorAddress.Hex())
		t.Logf("PROJECT_DID: %v", string(projectDid[:]))
		t.Logf("TRANSACTION_ID: %v", string(txID[:]))

		_, err := authContract.Validate(&validatorWallet, txID, projectWalletAuthAddress, projectWalletAddress, evaluatorAddress, big.NewInt(200000000))

		if err != nil {
			t.Errorf("Error: %v", err)
		}

	}
}
