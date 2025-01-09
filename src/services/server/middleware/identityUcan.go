package middleware

import (
	"context"
	"net/http"
	"storage/src/pkg/goucan"
	"storage/src/pkg/logger"
	identitymodule "storage/src/services/contracts/IdentityModule"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IdentityUcanVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.GetContextLogger(c)
		contractAddress := c.GetString("contract")

		if contractAddress == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Contract address not found"})
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		isAuthenticated, err := validateIdentityUcanPayload(c, contractAddress, token)
		if err != nil {
			log.Error("Error validating identity UCAN", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid UCAN", "error": err.Error()})
			c.Abort()
			return
		}

		log.Info("isAuthenticated", "isAuthenticated", isAuthenticated)

		c.Set("isAuthenticated", isAuthenticated)
		c.Next()
	}
}

func validateIdentityUcanPayload(c context.Context, contractAddress, token string) (bool, error) {
	log := logger.GetContextLogger(c)

	identityModuleContract, err := identitymodule.NewIdentityModuleContract(c, contractAddress)
	if err != nil {
		log.Error("Error creating identity module contract", "error", err)
		return false, err
	}

	identityDetails, err := identityModuleContract.Instance.GetIdentityModulePublicDetails(&bind.CallOpts{Context: c})
	if err != nil {
		log.Error("Error getting identity module public details", "error", err)
		return false, err
	}

	signingDid := identityDetails.SigningDid
	serviceDID := viper.GetString("service.did")
	payload := &goucan.UcanPayload{
		Iss: signingDid,
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
