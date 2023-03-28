package commonMethods

import (
	"crypto/sha256"
	"encoding/hex"
)

func StringToSHA256(str string) string {
	if str == "" {
		return ""
	}
	hasher := sha256.New()
	hasher.Write([]byte(str))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
