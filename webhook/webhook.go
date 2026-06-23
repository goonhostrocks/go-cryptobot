package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

var ErrInvalidSignature = errors.New("invalid webhook signature")

func VerifyUpdate(bodyRaw []byte, token string, receivedSignature string) bool {
	if receivedSignature == "" {
		return false
	}

	secretHash := sha256.Sum256([]byte(token))

	h := hmac.New(sha256.New, secretHash[:])
	h.Write(bodyRaw)

	expectedSignature := hex.EncodeToString(h.Sum(nil))

	return hmac.Equal([]byte(expectedSignature), []byte(receivedSignature))
}

func ParseAndVerify(bodyRaw []byte, token string, receivedSignature string) (*Update, error) {
	// 1. Authenticate the payload
	if !VerifyUpdate(bodyRaw, token, receivedSignature) {
		return nil, ErrInvalidSignature
	}

	// 2. Unmarshal into your exact structure
	var up Update
	if err := json.Unmarshal(bodyRaw, &up); err != nil {
		return nil, err
	}

	return &up, nil
}
