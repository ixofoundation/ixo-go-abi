package test

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	auth "github.com/ixofoundation/ixo-go-abi/abi/auth"
	project "github.com/ixofoundation/ixo-go-abi/abi/project"
	token "github.com/ixofoundation/ixo-go-abi/abi/token"
)

func TestAuthContract(t *testing.T) {
	// Owner Address
	keyAuth, _ := crypto.HexToECDSA("b6ad4d7b59a2766e94f9290740fd62676165684500c6d1331185912600e19481")
	owner := bind.NewKeyedTransactor(keyAuth)
	owner.GasLimit = uint64(2000000)
	owner.GasPrice = big.NewInt(2000000)

	// ValidatorNode Address
	keyVal, _ := crypto.HexToECDSA("b91ff5ab275b21e12482bca1481fa9726c27c9debda25c4c3c8cd192468d730e")
	validatorNode := bind.NewKeyedTransactor(keyVal)
	validatorNode.GasLimit = uint64(4500000)

	// Evaluator Address
	keyEval, _ := crypto.HexToECDSA("baf7af00a5f868db6ef8ca22ebbf69c131217ef08427c040d90931a803d98957")
	evalWallet := bind.NewKeyedTransactor(keyEval)

	// Use Ganache
	blockchain, _ := ethclient.Dial("http://127.0.0.1:7545")

	callOpts := bind.CallOpts{
		Pending: false,
		From:    owner.From,
		Context: nil,
	}

	// #1 ERC20
	ixoTokenAddress, _, ixoTokenContact, _ := token.DeployIxoERC20Token(
		owner,
		blockchain,
	)
	t.Logf("ERC20_ADDRESS: %v", ixoTokenAddress.Hex())

	// #2 AUTH
	var members []common.Address
	members = append(members, validatorNode.From)

	authContractAddress, _, authContract, _ := auth.DeployAuthContract(
		owner,
		blockchain,
		members,
		big.NewInt(1),
	)
	t.Logf("AUTH_ADDRESS: %v", authContractAddress.Hex())

	// #3 PROJECT_WALLET_ATH
	projectWalletAuthAddress, _, projectWalletAuthContract, _ := auth.DeployProjectWalletAuthoriser(
		owner,
		blockchain,
	)
	t.Logf("PROJECT_WALLET_AUTH_ADDRESS: %v", projectWalletAuthAddress.Hex())

	// #4 SET_AUTH
	projectWalletAuthContract.SetAuthoriser(owner, authContractAddress)

	// #5 PROJECT_WALLET_FACTORY
	projectWalletFactoryAddress, _, _, _ := project.DeployProjectWalletFactory(owner, blockchain)
	t.Logf("PROJECT_WALLET_FACTORY_ADDRESS: %v", projectWalletFactoryAddress.Hex())

	// #6 PROJECT_WALLET_REGISTRY
	projectWalletRegistryAddress, _, projectWalletRegistryContract, _ := project.DeployProjectWalletRegistry(owner, blockchain, ixoTokenAddress, authContractAddress, projectWalletFactoryAddress)
	t.Logf("PROJECT_WALLET_REGISTRY_ADDRESS: %v", projectWalletRegistryAddress.Hex())

	// #7 CREATE_PROJECT_DID
	var projectDid [32]byte
	copy(projectDid[:], "WjU6gE1JhZANcdv3aC8PEJ")

	// #8 CREATE_PROJECT_WALLET
	projectWalletRegistryContract.EnsureWallet(owner, projectDid)
	projectWalletAddress, _ := projectWalletRegistryContract.WalletOf(&callOpts, projectDid)
	t.Logf("PROJECT_WALLET_ADDRESS: %v", projectWalletAddress.Hex())

	// #9 SET_TOKEN_MINTER
	ixoTokenContact.SetMinter(owner, owner.From)

	// #10 FUND_PROJECT_WALLET
	ixoTokenContact.Mint(owner, projectWalletAddress, big.NewInt(500000000))

	// #11 CREATE_RANDOM_TX_ID
	var txID [32]byte
	copy(txID[:], randomString(32)[:32])

	// #12 RUN_UNIT_TESTS
	t.Run("Check if address is part of members list", checkIsMember(*authContract, callOpts, owner.From))
	t.Run("Check quorum size", checkQuorumSize(*authContract, callOpts))
	t.Run("Check member size", checkMemberCount(*authContract, callOpts))
	t.Run("Validate transaction", validateTransaction(*authContract, *validatorNode, txID, projectWalletAuthAddress, projectWalletAddress, evalWallet.From))
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(len int) []byte {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return bytes
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

func validateTransaction(authContract auth.AuthContract, transOpts bind.TransactOpts, transaction [32]byte, projectWalletAuthAddress common.Address, projectWalletAddress common.Address, evalAddress common.Address) func(*testing.T) {
	return func(t *testing.T) {

		var txn []byte = transaction[:]
		// transaction = Transaction hash
		// targetAddr = ProjectWalletAuthAddress
		// senderAddr = ProjectWalletAddress
		// receiverAddr = Person receiving funds
		t.Logf("validatorNode: %v", transOpts.From.Hex())
		t.Logf("transaction: %v", common.ToHex(txn))
		t.Logf("senderAddr: %v", projectWalletAddress.Hex())
		t.Logf("receiverAddr: %v", evalAddress.Hex())
		t.Logf("authAddr: %v", projectWalletAuthAddress.Hex())

		transaction, err := authContract.Validate(&transOpts, transaction, projectWalletAuthAddress, projectWalletAddress, evalAddress, big.NewInt(200000000))

		if err != nil {
			t.Errorf("ERROR: %v", err)
		}

		t.Logf("Transaction Hash: %v", transaction)
	}
}
