package identitymodule

import (
	"context"
	"storage/src/pkg/logger"
	"storage/src/services/contracts"
	identitymoduleabi "storage/src/services/contracts/IdentityModule/abi"

	"github.com/spf13/viper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IdentityModuleContract struct {
	Instance *identitymoduleabi.BindingsCaller
}

func NewIdentityModuleContract(ctx context.Context, contract string) (*IdentityModuleContract, error) {
	log := logger.GetContextLogger(ctx)

	network := viper.GetString("service.network")
	networkProviderUrl := contracts.NetworkProviderUrl(network)
	log.Info("network provider url", "url", networkProviderUrl)
	contractAddress := common.HexToAddress(contract)

	client, err := ethclient.Dial(networkProviderUrl)
	if err != nil {
		log.Error("Failed to connect to the Ethereum client", "error", err)
		return nil, err
	}

	contractInstance, err := identitymoduleabi.NewBindingsCaller(contractAddress, client)
	if err != nil {
		log.Error("Failed to instantiate a smart contract", "error", err)
		return nil, err
	}

	return &IdentityModuleContract{Instance: contractInstance}, nil
}
