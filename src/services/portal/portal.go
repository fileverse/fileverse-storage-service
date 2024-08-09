package portal

import (
	"context"
	"errors"
	"fmt"
	"storage/src/pkg/logger"
	"storage/src/services/portal/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PortalContract struct {
	Instance *abi.BindingsCaller
}

func NewPortalContract(ctx context.Context, contract, network, chainId string) (*PortalContract, error) {
	log := logger.GetContextLogger(ctx)
	if chainId == "" && network == "" {
		return nil, errors.New("atleast one of network or chainId is required")
	}

	if network == "" {
		network = networkFromChainId(chainId)
	}
	
	networkProviderUrl := networkProviderUrl(network)
	log.Info("network provider url", "url", networkProviderUrl)
	contractAddress := common.HexToAddress(contract)

	client, err := ethclient.Dial(networkProviderUrl)
	if err != nil {
		// log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return nil, err
	}

	contractInstance, err := abi.NewBindingsCaller(contractAddress, client)
	if err != nil {
		// log.Fatalf("Failed to instantiate a smart contract: %v", err)
		fmt.Printf("Failed to instantiate a smart contract: %v", err)
		return nil, err
	}

	return &PortalContract{contractInstance}, nil
}
