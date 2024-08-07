package goucan

import (
	"fmt"
	"strings"
)

type Capabilities struct {
	With interface{} `json:"with"`
	Can  interface{} `json:"can"`
}

func (cap *Capabilities) getReadableCap() (*Capabilities, error) {
	newCap := &Capabilities{}
	switch cap.With.(type) {
	case string:
		with := cap.With.(string)
		withParts := strings.Split(with, ":")

		newCap.With = map[string]interface{}{
			"schema":   withParts[0],
			"hierPart": withParts[1],
		}
	default:
		return nil, fmt.Errorf("invalid WITH type")
	}

	switch cap.Can.(type) {
	case string:
		can := cap.Can.(string)
		canParts := strings.Split(can, "/")

		newCap.Can = map[string]interface{}{
			"namespace": canParts[0],
			"segments":  canParts[1:],
		}

	default:
		return nil, nil
	}

	return newCap, nil
}
