package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
)

func TestIxoErc20TokenContract(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	blockchain := backends.NewSimulatedBackend(core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}}, uint64(10000000))

	_, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		auth,
		blockchain,
	)

	blockchain.Commit()

	callOpts := bind.CallOpts{
		Pending: false,
		From:    auth.From,
		Context: nil,
	}

	transOpts := bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}

	// Test Token Name
	t.Run("Check token name is correct", checkTokenName(*ixoTokenContact, callOpts))

	// Test Token Symbol
	t.Run("Check token symbol is correct", checkTokenSymbol(*ixoTokenContact, callOpts))

	// Test Token CAP
	t.Run("Check token cap is correct", checkTokenCap(*ixoTokenContact, callOpts))

	// Test Set Minter
	ixoTokenContact.SetMinter(&transOpts, auth.From)
	blockchain.Commit()
	t.Run("Check setting minter address", checkMinterAddress(*ixoTokenContact, callOpts, auth.From))

	// Test Mint Token to address
	amountToMint := big.NewInt(22000)
	ixoTokenContact.Mint(&transOpts, auth.From, amountToMint)
	blockchain.Commit()
	t.Run("Check minting amount to address", mintTokenToAddress(*ixoTokenContact, callOpts, auth.From, *amountToMint))
}

func checkTokenName(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		if tokenName, _ := ixoTokenContact.Name(&callOpts); tokenName != "IXO Token" {
			t.Errorf("Expected tokenName to be: IXO Token")
		}
	}
}

func checkTokenSymbol(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		if tokenSymbol, _ := ixoTokenContact.Symbol(&callOpts); tokenSymbol != "IXO" {
			t.Errorf("Expected tokenName to be: IXO")
		}
	}
}

func checkTokenCap(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		tokenCapCompare := big.NewInt(1000000000000000000)
		if tokenCap, _ := ixoTokenContact.CAP(&callOpts); tokenCap.Cmp(tokenCapCompare) != 0 {
			t.Errorf("Expected tokenCap to be: %v", tokenCapCompare)
			t.Errorf("But got tokenCap: %v", tokenCap)
		}
	}
}

func checkMinterAddress(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts, minterAddress common.Address) func(*testing.T) {
	return func(t *testing.T) {
		if minter, _ := ixoTokenContact.Minter(&callOpts); minterAddress != minter {
			t.Errorf("Expected minter to be: %v", minterAddress)
			t.Errorf("But got minter: %v", minter)
		}
	}
}

func mintTokenToAddress(ixoTokenContact token.IxoERC20Token, callOpts bind.CallOpts, minterAddress common.Address, tokensMinted big.Int) func(*testing.T) {
	return func(t *testing.T) {
		if tokenBalance, _ := ixoTokenContact.BalanceOf(&callOpts, minterAddress); tokensMinted.Cmp(tokenBalance) != 0 {
			t.Errorf("Expected token balance to be: %v", tokensMinted)
		}
	}
}
