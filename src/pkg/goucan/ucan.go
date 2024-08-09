package goucan

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"reflect"
	"storage/src/pkg/logger"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mr-tron/base58"
)

type Ucan struct {
	Header     UcanHeader
	Payload    UcanPayload
	DataToSign []byte
	Signature  []byte
}

type UcanHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	Ucv       string `json:"ucv"`
}

type UcanPayload struct {
	Iss string         `json:"iss"`
	Aud string         `json:"aud"`
	Exp int64          `json:"exp"`
	Att []Capabilities `json:"att"`
	Prf []string       `json:"prf"`
}

func (uc *Ucan) extractPublicKeyBytesFromDID(did string) ([]byte, error) {
	// Check if the DID starts with the correct prefix
	if !strings.HasPrefix(did, BASE58_DID_PREFIX) {
		return nil, ErrDidPrefixNotFound
	}

	// Remove the prefix
	base58Key := strings.TrimPrefix(did, BASE58_DID_PREFIX)

	// Decode the base58 string
	decodedBytes, _ := base58.Decode(base58Key)

	// Check if the decoded bytes start with the Edwards prefix
	if len(decodedBytes) < len(EDWARDS_DID_PREFIX) || string(decodedBytes[:len(EDWARDS_DID_PREFIX)]) != EDWARDS_DID_PREFIX {
		return nil, ErrEdwardsPrefixNotFound
	}

	// Remove the Edwards prefix to get the public key bytes
	publicKeyBytes := decodedBytes[len(EDWARDS_DID_PREFIX):]
	return publicKeyBytes, nil
}

func (uc *Ucan) checkEdDsaSignature(c context.Context) error {
	publicKey, err := uc.extractPublicKeyBytesFromDID(uc.Payload.Iss)
	if err != nil {
		return fmt.Errorf("error extracting public key bytes from DID: %w", err)
	}

	signer := &jwt.SigningMethodEd25519{}
	signingkey := ed25519.PublicKey(publicKey)
	signingData := string(uc.DataToSign)
	signature := base64.RawURLEncoding.EncodeToString(uc.Signature)
	err = signer.Verify(signingData, signature, signingkey)
	return err
}

func (uc *Ucan) isExpired(ctx context.Context) bool {
	log := logger.GetContextLogger(ctx)
	if uc.Payload.Exp == 0 {
		return false
	}

	log.Debug("Payload.Exp", "Payload.Exp", uc.Payload.Exp)
	log.Debug("TimeFunc", "TimeFunc", jwt.TimeFunc().Unix())
	return uc.Payload.Exp < jwt.TimeFunc().Unix()
}

func (uc *Ucan) getReadableAtt() ([]Capabilities, error) {
	att := make([]Capabilities, len(uc.Payload.Att))
	for i, cap := range uc.Payload.Att {
		newCap, err := cap.getReadableCap()
		if err != nil {
			return nil, err
		}
		att[i] = *newCap
	}
	return att, nil
}

func (uc *Ucan) compareAtt(ctx context.Context, att []Capabilities) bool {
	log := logger.GetContextLogger(ctx)
	ucAtt, err := uc.getReadableAtt()
	if len(ucAtt) != len(att) || err != nil {
		log.Error("Error getting readable att", "error", err)
		return false
	}

	capMap1 := make(map[string]bool)
	capMap2 := make(map[string]bool)

	for _, cap := range ucAtt {
		capMap1[fmt.Sprintf("%v:%v", cap.With, cap.Can)] = true
	}

	for _, cap := range att {
		capMap2[fmt.Sprintf("%v:%v", cap.With, cap.Can)] = true
	}

	return reflect.DeepEqual(capMap1, capMap2)
}

func (uc *Ucan) Verify(ctx context.Context, payload *UcanPayload) (bool, error) {
	err := uc.checkEdDsaSignature(ctx)
	if err != nil {
		return false, ErrUcanSignatureInvalid
	}

	if uc.isExpired(ctx) {
		return false, ErrUcanExpired
	}

	if uc.Payload.Aud != payload.Aud {
		return false, fmt.Errorf("aud mismatch")
	}

	if uc.Payload.Iss != payload.Iss {
		return false, fmt.Errorf("iss mismatch")
	}

	if !uc.compareAtt(ctx, payload.Att) {
		return false, fmt.Errorf("att mismatch")
	}

	return true, nil
}
