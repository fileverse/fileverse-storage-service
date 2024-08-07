package constants

import "os"

var (
	ETH_GOERLI_RPC_URL      string = os.Getenv("ETH_GOERLI_RPC_URL")
	ETH_MAINNET_RPC_URL     string = os.Getenv("ETH_MAINNET_RPC_URL")
	ETH_SEPOLIA_RPC_URL     string = os.Getenv("ETH_SEPOLIA_RPC_URL")
	GNOSIS_MAINNET_RPC_URL  string = os.Getenv("GNOSIS_MAINNET_RPC_URL")
	GNOSIS_TESTNET_RPC_URL  string = os.Getenv("GNOSIS_TESTNET_RPC_URL")
	POLYGON_MAINNET_RPC_URL string = os.Getenv("POLYGON_MAINNET_RPC_URL")
)
