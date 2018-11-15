package test

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
	"github.com/stretchr/testify/assert"
)

func TestBasicProjectWalletContract(t *testing.T) {
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

	// #1 ERC20
	ixoTokenAddress, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		ownerWallet,
		blockchain,
	)
	t.Logf("ERC20_ADDRESS: %v", ixoTokenAddress.Hex())

	// SET_MINTER
	ixoTokenContact.SetMinter(ownerWallet, ownerWallet.From)

	basicProjectWalletAddress, _, basicProjectWalletContract, _ := project.DeployBasicProjectWallet(
		ownerWallet,
		blockchain,
		ixoTokenAddress,
		ownerWallet.From,
		random32Bytes(),
	)
	t.Logf("BASIC_PROJECT_WALLET_ADDRESS: %v", basicProjectWalletAddress.Hex())

	// MINT_TOKENS_TO_PROJECT_WALLET
	ixoTokenContact.Mint(ownerWallet, basicProjectWalletAddress, big.NewInt(800000000))

	t.Run("Check balance on basic project wallet", checkBalanceOnBasicProjectWallet(*ixoTokenContact, callOpts, basicProjectWalletAddress, *big.NewInt(800000000)))
	t.Run("Transfer tokens from basic project wallet to evaluatorWallet", transferTokensFromBasicProjectWallet(*basicProjectWalletContract, *ixoTokenContact, callOpts, *ownerWallet, evaluatorWallet.From, *big.NewInt(400000000)))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func random32Bytes() [32]byte {
	bytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	var txID [32]byte
	copy(txID[:], bytes[:32])
	return txID
}

func checkBalanceOnBasicProjectWallet(tokenContract token.IxoERC20Token, callOpts bind.CallOpts, basicProjectWalletAddress common.Address, expectedBalance big.Int) func(*testing.T) {
	return func(t *testing.T) {
		balance, _ := tokenContract.BalanceOf(&callOpts, basicProjectWalletAddress)
		assert.EqualValues(t, &expectedBalance, balance, "Incorrect balance!")
	}
}

func transferTokensFromBasicProjectWallet(basicProjectWalletContract project.BasicProjectWallet, tokenContract token.IxoERC20Token, callOpts bind.CallOpts, transOpts bind.TransactOpts, evaluatorAddress common.Address, tokenAmount big.Int) func(*testing.T) {
	return func(t *testing.T) {
		basicProjectWalletContract.Transfer(&transOpts, evaluatorAddress, &tokenAmount)
		evaluatorBalance, _ := tokenContract.BalanceOf(&callOpts, evaluatorAddress)
		assert.EqualValues(t, &tokenAmount, evaluatorBalance, "Incorrect balance!")
	}
}