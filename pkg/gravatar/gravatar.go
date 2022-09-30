package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateHash(email string) string {
	norm := strings.ToLower(strings.Trim(email, " "))
	hash := md5.Sum([]byte(norm))
	return hex.EncodeToString(hash[:])
}

func BuildUrl(hash string) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%v?d=retro&f=y", hash)
}
