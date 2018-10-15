package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	abi "github.com/ixofoundation/ixo-go-abi/abi"
)

func TestGetTokenCap(t *testing.T) {

	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 80000000)

	_, _, ixoTokenContact, _ := abi.DeployIxoERC20Token(
		auth,
		blockchain,
	)

	blockchain.Commit()

	callOpts := bind.CallOpts{
		Pending: false,
		From:    auth.From,
		Context: nil,
	}

	tokenCapCompare := big.NewInt(1000000000000000000)

	if tokenCap, _ := ixoTokenContact.CAP(&callOpts); tokenCap.Cmp(tokenCapCompare) != 0 {
		t.Errorf("Expected tokenCap to be: %v", tokenCapCompare)
		t.Errorf("But got tokenCap: %v", tokenCap)
	}
}
