package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
	util "github.com/ixofoundation/ixo-go-abi/test/util"
	"github.com/stretchr/testify/assert"
)

func TestIxoErc20TokenContract(t *testing.T) {
	// Use Ganache Account
	key, _ := crypto.HexToECDSA("b6ad4d7b59a2766e94f9290740fd62676165684500c6d1331185912600e19481")
	authorizer := bind.NewKeyedTransactor(key)

	// Use Ganache
	blockchain, _ := ethclient.Dial("http://127.0.0.1:7545")

	_, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		authorizer,
		blockchain,
	)

	callOpts := bind.CallOpts{
		Pending: false,
		From:    authorizer.From,
		Context: nil,
	}

	transOpts := bind.TransactOpts{
		From:   authorizer.From,
		Signer: authorizer.Signer,
	}

	// Test Set Minter
	ixoTokenContact.SetMinter(&transOpts, authorizer.From)

	// Test Mint Token to address
	amountToMint := big.NewInt(22000)
	ixoTokenContact.Mint(&transOpts, authorizer.From, amountToMint)

	t.Run("Check token name is correct", checkTokenName(*ixoTokenContact, callOpts))
	t.Run("Check token symbol is correct", checkTokenSymbol(*ixoTokenContact, callOpts))
	t.Run("Check token cap is correct", checkTokenCap(*ixoTokenContact, callOpts))
	t.Run("Check setting minter address", checkMinterAddress(*ixoTokenContact, callOpts, authorizer.From))
	t.Run("Check minting amount to address", mintTokenToAddress(*ixoTokenContact, callOpts, authorizer.From, *amountToMint))
}

func checkTokenName(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		tokenName, _ := ixoTokenContact.Name(&callOpts)
		assert.Equal(t, "IXO Token", tokenName, "Incorrect token name!")
	}
}

func checkTokenSymbol(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		tokenSymbol, _ := ixoTokenContact.Symbol(&callOpts)
		assert.Equal(t, "IXO", tokenSymbol, "Incorrect token symbol!")
	}
}

func checkTokenCap(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		tokenCap, _ := ixoTokenContact.CAP(&callOpts)
		assert.Equal(t, big.NewInt(1000000000000000000), tokenCap, "Incorrect token cap!")
	}
}

func checkMinterAddress(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts, minterAddress common.Address) func(*testing.T) {
	return func(t *testing.T) {
		minter, _ := ixoTokenContact.Minter(&callOpts)
		assert.Equal(t, minterAddress, minter, "Incorrect minter!")
	}
}

func mintTokenToAddress(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts, minterAddress common.Address, tokensMinted big.Int) func(*testing.T) {
	return func(t *testing.T) {
		util.CheckBalance(ixoTokenContact, callOpts, minterAddress, tokensMinted)
	}
}
