package middleware

import (
	"context"
	"storage/src/pkg/goucan"
	"storage/src/pkg/logger"
)

// validateUcanPayload validates a UCAN payload
func validateUcanPayload(c context.Context, payload *goucan.UcanPayload, token string) (bool, error) {
	log := logger.GetContextLogger(c)
	ucan, err := goucan.DecodeUcanString(c, token)
	if err != nil {
		log.Debug("Error decoding UCAN", "error", err)
		return false, err
	}

	result, err := ucan.Verify(c, payload)
	if err != nil {
		log.Debug("Error verifying UCAN", "error", err)
		return false, err
	}

	return result, nil
}
