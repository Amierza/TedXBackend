package helpers

import (
	"crypto/sha512"
	"encoding/hex"
)

func GenerateSignature(orderID, statusCode, grossAmount, serverKey string) string {
	raw := orderID + statusCode + grossAmount + serverKey
	hash := sha512.Sum512([]byte(raw))
	return hex.EncodeToString(hash[:])
}
