package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	auth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
)

func TestIxoErc20TokenContract(t *testing.T) {
	key, _ := crypto.GenerateKey()
	authorizer := bind.NewKeyedTransactor(key)
	projectWallet := bind.NewKeyedTransactor(key)
	blockchain := backends.NewSimulatedBackend(core.GenesisAlloc{authorizer.From: {Balance: big.NewInt(10000000000)}}, uint64(10000000))

	callOpts := bind.CallOpts{
		Pending: false,
		From:    authorizer.From,
		Context: nil,
	}

	transOpts := bind.TransactOpts{
		From:   authorizer.From,
		Signer: authorizer.Signer,
		GasLimit: uint64(100000),
	}

	// Setup quorum
	var quorum []common.Address
	quorum = append(quorum, authorizer.From)

	// Deploy auth contract
	_, _, authContract, _ := auth.DeployAuthContract(
		authorizer,
		blockchain,
		quorum,
		big.NewInt(1),
	)

	_, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		authorizer,
		blockchain,
	)

	projectWalletAuthAddress, _, _, _ := auth.DeployProjectWalletAuthoriser(
		authorizer,
		blockchain,
	)

	blockchain.Commit()

	// Simulate token mint and transfer

	ixoTokenContact.Mint(&transOpts, authorizer.From, big.NewInt(50000))
	blockchain.Commit()

	transaction, err := ixoTokenContact.TransferFrom(&transOpts, projectWallet.From, authorizer.From, big.NewInt(50000))

	if err != nil {
		t.Errorf("ERROR: %v", err)
	}

	t.Run("Check if address is part of members list", checkIsMember(*authContract, callOpts, authorizer.From))
	t.Run("Check quorum size", checkQuorumSize(*authContract, callOpts))
	t.Run("Check member size", checkMemberCount(*authContract, callOpts))
	t.Run("Validate transaction", validateTransaction(*authContract, transOpts, transaction.Hash(), projectWallet.From, authorizer.From, projectWalletAuthAddress))
}

func checkIsMember(authContract auth.AuthContract, callOpts bind.CallOpts, address common.Address) func(*testing.T) {
	return func(t *testing.T) {
		if isMember, _ := authContract.IsMember(&callOpts, address); !isMember {
			t.Errorf("Address is not part of members list")
		}
	}
}

func checkQuorumSize(authContract auth.AuthContract, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		quorumSize, _ := authContract.Quorum(&callOpts)

		if quorumSize.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Quorum needs to be size: %v", big.NewInt(1))
		}
	}
}

func checkMemberCount(authContract auth.AuthContract, callOpts bind.CallOpts) func(*testing.T) {
	return func(t *testing.T) {
		memberCount, _ := authContract.MemberCount(&callOpts)

		if memberCount.Cmp(big.NewInt(1)) != 0 {
			t.Errorf("Member count needs to be: %v", big.NewInt(1))
		}
	}
}

func validateTransaction(authContract auth.AuthContract, transOpts bind.TransactOpts, transaction [32]byte, senderAddr common.Address, receiverAddr common.Address, targetAddr common.Address) func(*testing.T) {
	return func(t *testing.T) {

		// transaction = Transaction hash
		// targetAddr = ProjectWalletAuthAddress
		// senderAddr = ProjectWalletAddress
		// receiverAddr = Person receiving funds
		transaction, err := authContract.Validate(&transOpts, transaction, targetAddr, senderAddr, receiverAddr, big.NewInt(20))
		if err != nil {
			t.Errorf("ERROR: %v", err)
		}

		t.Logf("Transaction Hash: %v", transaction.Hash())
	}
}
