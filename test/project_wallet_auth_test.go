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

func TestProjectWalletAuthContract(t *testing.T) {
	// Contract Owner Address
	keyAuth, _ := crypto.HexToECDSA("b6ad4d7b59a2766e94f9290740fd62676165684500c6d1331185912600e19481")
	ownerWallet := bind.NewKeyedTransactor(keyAuth)
	ownerWallet.GasLimit = uint64(2782100)

	// Evaluator Address
	keyEval, _ := crypto.HexToECDSA("baf7af00a5f868db6ef8ca22ebbf69c131217ef08427c040d90931a803d98957")
	evaluatorWallet := bind.NewKeyedTransactor(keyEval)

	// Use Ganache
	blockchain, _ := ethclient.Dial("http://127.0.0.1:7545")

	callOpts := bind.CallOpts{
		Pending: false,
		From:    ownerWallet.From,
		Context: nil,
	}

	ixoTokenAddress, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		ownerWallet,
		blockchain,
	)
	t.Logf("ERC20_ADDRESS: %v", ixoTokenAddress.Hex())

	// SET_MINTER
	ixoTokenContact.SetMinter(ownerWallet, ownerWallet.From)

	projectWalletAuthAddress, _, projectWalletAuthContract, _ := auth.DeployProjectWalletAuthoriser(ownerWallet, blockchain)

	t.Logf("PROJECT_WALLET_AUTH_ADDRESS: %v", projectWalletAuthAddress.Hex())

	// SET_PROJECT_WALLET_AUTH_OWNER
	projectWalletAuthContract.SetAuthoriser(ownerWallet, ownerWallet.From)

	basicProjectWalletAddress, _, _, _ := project.DeployBasicProjectWallet(
		ownerWallet,
		blockchain,
		ixoTokenAddress,
		projectWalletAuthAddress,
		util.Random32Bytes(),
	)
	t.Logf("BASIC_PROJECT_WALLET_ADDRESS: %v", basicProjectWalletAddress.Hex())

	// MINT_TOKENS_TO_PROJECT_WALLET
	ixoTokenContact.Mint(ownerWallet, basicProjectWalletAddress, big.NewInt(800000000))

	t.Run("Check balance on basic project wallet", util.CheckBalance(*ixoTokenContact, callOpts, basicProjectWalletAddress, *big.NewInt(800000000)))
	t.Run("Transfer tokens from basic project wallet to evaluatorWallet", transferTokensUsingProjectWalletAuth(*projectWalletAuthContract, *ixoTokenContact, callOpts, *ownerWallet, basicProjectWalletAddress, evaluatorWallet.From, *big.NewInt(400000000)))
}

func transferTokensUsingProjectWalletAuth(projectWalletAuthContract auth.ProjectWalletAuthoriser, tokenContract token.IxoERC20Token, callOpts bind.CallOpts, transOpts bind.TransactOpts, basicProjectWalletAddress common.Address, evaluatorAddress common.Address, tokenAmount big.Int) func(*testing.T) {
	return func(t *testing.T) {
		projectWalletAuthContract.Transfer(&transOpts, basicProjectWalletAddress, evaluatorAddress, &tokenAmount)
		evaluatorBalance, _ := tokenContract.BalanceOf(&callOpts, evaluatorAddress)
		assert.EqualValues(t, &tokenAmount, evaluatorBalance, "Incorrect balance!")
	}
}
