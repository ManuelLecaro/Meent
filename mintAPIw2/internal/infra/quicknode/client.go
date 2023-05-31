package quicknode

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"mintapi/internal/infra/config"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthBaseClient struct {
	Client  *ethclient.Client
	Configs *config.GlobalConfigurations
}

// NewEthBaseClient creates a new instance of the EthBaseClient
func NewEthBaseClient(dir string) *EthBaseClient {
	client, err := ethclient.Dial(dir)
	if err != nil {
		return nil
	}

	return &EthBaseClient{
		Client:  client,
		Configs: config.GetGlobalConfigs(),
	}
}

func (ec *EthBaseClient) GetOwnerData(client *ethclient.Client, privateKeyStr string, value, gasLimit int64) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value)   // in wei
	auth.GasLimit = uint64(gasLimit) // 300000 in units
	auth.GasPrice = gasPrice

	return auth
}
