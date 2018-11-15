package util

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
	"github.com/stretchr/testify/assert"
)

// Random32Bytes generates random [32]byte
func Random32Bytes() [32]byte {
	key, _ := crypto.GenerateKey()
	randomTxID := bind.NewKeyedTransactor(key)
	randomTxID.From.Hex()
	var txID [32]byte
	copy(txID[:], randomTxID.From.Hex()[:32])

	return txID
}

// CheckBalance asserts the balance on a wallet
func CheckBalance(tokenContract token.IxoERC20Token, callOpts bind.CallOpts, walletAddress common.Address, expectedBalance big.Int) func(*testing.T) {
	return func(t *testing.T) {
		balance, _ := tokenContract.BalanceOf(&callOpts, walletAddress)
		assert.EqualValues(t, &expectedBalance, balance, "Incorrect balance!")
	}
}
