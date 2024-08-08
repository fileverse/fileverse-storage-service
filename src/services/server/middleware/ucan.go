package middleware

import (
	"context"
	goucan "storage/src/pkg/goucan"
	"storage/src/pkg/logger"
	"storage/src/services/portal"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Config represents the configuration for the service
var serviceDID string

func init() {
	serviceDID = viper.GetString("service.did")
}

// Verify is a Gin middleware for authentication and authorization
func UcanVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		contractAddress := c.GetString("contract")
		invokerAddress := c.GetString("invoker")
		chainID := c.GetString("chain")
		namespace := c.GetString("namespace")

		token := c.GetHeader("Authorization")
		if token == "" || invokerAddress == "" {
			c.Next()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		var err error
		var isAuthenticated bool

		switch {
		case namespace != "":
			isAuthenticated, err = validateNamespace(c, namespace, invokerAddress, token)
		case contractAddress != "":
			isAuthenticated, err = validateContractAddress(c, contractAddress, invokerAddress, token, chainID)
		default:
			isAuthenticated, err = validateInvokerAddress(c, invokerAddress, token)
		}

		if err != nil {
			log.Error("Error in verification", "error", err)
		}

		c.Set("isAuthenticated", isAuthenticated)
		c.Next()
	}
}

// validateUcanPayload validates a UCAN payload
func validateUcanPayload(c context.Context, payload *goucan.UcanPayload, token string) (bool, error) {
	ucan, err := goucan.DecodeUcanString(token)
	if err != nil {
		return false, err
	}

	result, err := ucan.Verify(payload)
	if err != nil {
		return false, err
	}

	return result, nil
}

// validateNamespace validates a namespace
func validateNamespace(c context.Context, namespace, invokerAddress, token string) (bool, error) {
	payload := &goucan.UcanPayload{
		Iss: invokerAddress,
		Aud: serviceDID,
		Att: []goucan.Capabilities{
			{
				With: map[string]interface{}{
					"schema":   "storage",
					"hierPart": invokerAddress,
				},
				Can: map[string]interface{}{
					"namespace": namespace,
					"segments":  []string{"CREATE", "GET"},
				},
			},
		},
	}

	return validateUcanPayload(c, payload, token)
}

// validateContractAddress validates a contract address
func validateContractAddress(c context.Context, contractAddress, invokerAddress, token, chainID string) (bool, error) {
	portal, err := portal.NewPortalContract(contractAddress, "", chainID)
	if err != nil {
		return false, err
	}

	keys, err := portal.Instance.CollaboratorKeys(&bind.CallOpts{Context: c}, common.HexToAddress(invokerAddress))
	if err != nil {
		return false, err
	}

	invokerDid := keys.EditDid

	payload := &goucan.UcanPayload{
		Iss: invokerDid,
		Aud: serviceDID,
		Att: []goucan.Capabilities{
			{
				With: map[string]interface{}{
					"schema":   "storage",
					"hierPart": strings.ToLower(contractAddress),
				},
				Can: map[string]interface{}{
					"namespace": "file",
					"segments":  []string{"CREATE"},
				},
			},
		},
	}

	return validateUcanPayload(c, payload, token)
}

// validateInvokerAddress validates an invoker address
func validateInvokerAddress(c context.Context, invokerAddress, token string) (bool, error) {
	payload := &goucan.UcanPayload{
		Iss: invokerAddress,
		Aud: serviceDID,
		Att: []goucan.Capabilities{
			{
				With: map[string]interface{}{
					"schema":   "storage",
					"hierPart": invokerAddress,
				},
				Can: map[string]interface{}{
					"namespace": "file",
					"segments":  []string{"CREATE", "GET"},
				},
			},
		},
	}

	return validateUcanPayload(c, payload, token)
}
