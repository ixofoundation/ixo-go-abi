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
)

func TestAuthContract(t *testing.T) {

	// Contract Owner Address
	keyAuth, _ := crypto.HexToECDSA("b6ad4d7b59a2766e94f9290740fd62676165684500c6d1331185912600e19481")
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

	callOpts := bind.CallOpts{
		Pending: false,
		From:    ownerWallet.From,
		Context: nil,
	}

	// DEPLOY_ERC20_CONTRACT
	ixoTokenAddress, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		ownerWallet,
		blockchain,
	)

	// SET_MINTER
	ixoTokenContact.SetMinter(ownerWallet, ownerWallet.From)

	// CREATE_MEMBERS_LIST
	members := []common.Address{validatorWallet.From}

	// DEPLOY_AUTH_CONTRACT
	authAddress, _, authContract, _ := auth.DeployAuthContract(ownerWallet, blockchain, members, big.NewInt(1))

	// DEPLOY_PROJECT_WALLET_AUTH_CONTRACT
	projectWalletAuthAddress, _, projectWalletAuthContract, _ := auth.DeployProjectWalletAuthoriser(ownerWallet, blockchain)

	// SET_PROJECT_WALLET_AUTH_OWNER
	projectWalletAuthContract.SetAuthoriser(ownerWallet, authAddress)

	// DEPLOY_BASIC_PROJECT_WALLET_CONTRACT
	basicProjectWalletAddress, _, _, _ := project.DeployBasicProjectWallet(
		ownerWallet,
		blockchain,
		ixoTokenAddress,
		projectWalletAuthAddress,
		util.Random32Bytes(),
	)

	// MINT_TOKENS_TO_PROJECT_WALLET
	ixoTokenContact.Mint(ownerWallet, basicProjectWalletAddress, big.NewInt(800000000))

	t.Run("Check balance on basic project wallet", util.CheckBalance(*ixoTokenContact, callOpts, basicProjectWalletAddress, *big.NewInt(800000000)))
	t.Run("Transfer tokens using auth contract", transferTokensUsingAuthContract(*authContract, *ixoTokenContact, callOpts, *validatorWallet, util.Random32Bytes(), projectWalletAuthAddress, basicProjectWalletAddress, evaluatorWallet.From, *big.NewInt(400000000)))
}

func transferTokensUsingAuthContract(authContract auth.AuthContract, tokenContract token.IxoERC20Token, callOpts bind.CallOpts, transOpts bind.TransactOpts, txID [32]byte, projectWalletAuthAddress common.Address, basicProjectWalletAddress common.Address, evaluatorAddress common.Address, tokenAmount big.Int) func(*testing.T) {
	return func(t *testing.T) {
		_, err := authContract.Validate(&transOpts, txID, projectWalletAuthAddress, basicProjectWalletAddress, evaluatorAddress, &tokenAmount)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		util.CheckBalance(tokenContract, callOpts, evaluatorAddress, tokenAmount)
	}
}
