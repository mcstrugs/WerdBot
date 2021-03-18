package util

import (
	"encoding/base32"
	"strings"
)

func HandleBase32(command string) string {
	words := strings.TrimPrefix(command, "!base32 ")
	encoded := base32.StdEncoding.EncodeToString([]byte(words))
	return encoded
}
