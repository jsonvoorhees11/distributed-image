package stringHelper

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func GetHashCode(input string, length ...int) string {
	h := sha1.New()
	io.WriteString(h, input)
	if len(length) > 0 {
		return hex.EncodeToString(h.Sum(nil)[0:length[0]])
	} else {
		return hex.EncodeToString(h.Sum(nil))
	}
}
