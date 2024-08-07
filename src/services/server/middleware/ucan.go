package middleware

import (
	"context"
	"log"
	goucan "storage/src/pkg/goucan"
	"storage/src/services/portal"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// Config represents the configuration for the service
var serviceDID string

func init() {
	serviceDID = viper.GetString("service.did")
}

// Verify is a Gin middleware for authentication and authorization
func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		contractAddress := c.GetHeader("Contract")
		invokerAddress := c.GetHeader("Invoker")
		chainID := c.GetHeader("Chain")
		namespace := c.GetHeader("Namespace")

		requestID := uuid.New().String()
		c.Set("requestId", requestID)
		c.Set("isAuthenticated", false)
		c.Set("invokerAddress", invokerAddress)
		c.Set("contractAddress", contractAddress)
		c.Set("chainId", chainID)
		c.Set("namespace", namespace)
		log.Printf("req.requestId: %s", requestID)

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
			log.Printf("Error in verification: %v", err)
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
