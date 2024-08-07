package goucan

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	ErrUcanSignatureInvalid  = fmt.Errorf("invalid signature")
	ErrEdwardsPrefixNotFound = fmt.Errorf("edwards prefix not found")
	ErrDidPrefixNotFound     = fmt.Errorf("dID prefix not found")
	ErrEncoding              = fmt.Errorf("invalid encoding")
	ErrUcanForamt            = fmt.Errorf("invalid Ucan foramt")
	ErrUcanExpired           = fmt.Errorf("ucan expired")
	ErrUcanNotActive         = fmt.Errorf("not active yet")
)

func DecodeUcanString(ucStr string) (*Ucan, error) {
	parts := strings.Split(ucStr, ".")
	if len(parts) != 3 {
		return nil, ErrUcanForamt
	}
	dataToSign := strings.Join(parts[:2], ".")

	var err error
	// var encoding mb.Encoding
	partsBytes := make([][]byte, 3)
	for i := range parts {
		partsBytes[i], err = base64.RawURLEncoding.DecodeString(parts[i])
		if err != nil {
			return nil, fmt.Errorf("%d-:-%v", ErrEncoding, err)
		}
	}

	fmt.Println(string(partsBytes[0]))
	fmt.Println(string(partsBytes[1]))

	header, err := DecodeUcanHeaderBytes(partsBytes[0])
	if err != nil {
		return nil, err
	}
	payload, err := DecodeUcanPayloadBytes(partsBytes[1])
	if err != nil {
		return nil, err
	}
	signature := partsBytes[2]

	return &Ucan{
		Header:     *header,
		Payload:    *payload,
		DataToSign: []byte(dataToSign),
		Signature:  signature,
	}, nil
}

func DecodeUcanHeaderBytes(uhBytes []byte) (*UcanHeader, error) {
	uh := &UcanHeader{}
	err := json.Unmarshal(uhBytes, uh)
	return uh, err
}

func DecodeUcanPayloadBytes(upBytes []byte) (*UcanPayload, error) {
	up := &UcanPayload{}
	err := json.Unmarshal(upBytes, up)
	return up, err
}
