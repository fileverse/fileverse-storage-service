package portal

import (
	"github.com/spf13/viper"
)

func networkFromChainId(chainId string) string {
	switch chainId {
	case "5":
		return "eth_goerli"
	case "11155111":
		return "eth_sepolia"
	case "8420":
		return "fileverse_testnet"
	case "1":
		return "eth_mainnet"
	case "137":
		return "polygon_mainnet"
	case "100":
		return "gnosis_mainnet"
	case "10200":
		return "gnosis_testnet"
	default:
		return "eth_goerli"
	}
}

func networkProviderUrl(network string) string {
	switch network {
	case "eth_goerli":
		return viper.GetString("rpc_url.eth_goerli")
	case "eth_sepolia":
		return viper.GetString("rpc_url.eth_sepolia")
	case "eth_mainnet":
		return viper.GetString("rpc_url.eth_mainnet")
	case "polygon_mainnet":
		return viper.GetString("rpc_url.polygon_mainnet")
	case "gnosis_mainnet":
		return viper.GetString("rpc_url.gnosis_mainnet")
	case "gnosis_testnet":
		return viper.GetString("rpc_url.gnosis_testnet")
	default:
		return ""
	}
}
