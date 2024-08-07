package constants

import (
	"github.com/spf13/viper"
)

var (
	ETH_GOERLI_RPC_URL      string = viper.GetString("rpc_url.eth_goerli")
	ETH_MAINNET_RPC_URL     string = viper.GetString("rpc_url.eth_mainnet")
	ETH_SEPOLIA_RPC_URL     string = viper.GetString("rpc_url.eth_sepolia")
	GNOSIS_MAINNET_RPC_URL  string = viper.GetString("rpc_url.gnosis_mainnet")
	GNOSIS_TESTNET_RPC_URL  string = viper.GetString("rpc_url.gnosis_testnet")
	POLYGON_MAINNET_RPC_URL string = viper.GetString("rpc_url.polygon_mainnet")
)
