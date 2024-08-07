package portal

import "storage/src/constants"

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
		return constants.ETH_GOERLI_RPC_URL
	case "eth_sepolia":
		return constants.ETH_SEPOLIA_RPC_URL
	case "eth_mainnet":
		return constants.ETH_MAINNET_RPC_URL
	case "polygon_mainnet":
		return constants.POLYGON_MAINNET_RPC_URL
	case "gnosis_mainnet":
		return constants.GNOSIS_MAINNET_RPC_URL
	case "gnosis_testnet":
		return constants.GNOSIS_TESTNET_RPC_URL
	default:
		return ""
	}
}
