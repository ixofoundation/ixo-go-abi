package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
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

	// Test Token Name
	t.Run("Check token name is correct", checkTokenName(*ixoTokenContact, callOpts))

	// Test Token Symbol
	t.Run("Check token symbol is correct", checkTokenSymbol(*ixoTokenContact, callOpts))

	// Test Token CAP
	t.Run("Check token cap is correct", checkTokenCap(*ixoTokenContact, callOpts))

	// Test Set Minter
	ixoTokenContact.SetMinter(&transOpts, authorizer.From)
	t.Run("Check setting minter address", checkMinterAddress(*ixoTokenContact, callOpts, authorizer.From))

	// Test Mint Token to address
	amountToMint := big.NewInt(22000)
	ixoTokenContact.Mint(&transOpts, authorizer.From, amountToMint)
	t.Run("Check minting amount to address", mintTokenToAddress(*ixoTokenContact, callOpts, authorizer.From, *amountToMint))
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
