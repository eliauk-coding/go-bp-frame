package crypto

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Sha1Sum(raw string) string {
	s := sha1.New()
	io.WriteString(s, raw)
	return fmt.Sprintf("%x", s.Sum(nil))
}
