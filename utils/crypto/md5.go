package crypto

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5Sum(raw string) string {
	m := md5.New()
	io.WriteString(m, raw)
	return fmt.Sprintf("%x", m.Sum(nil))
}
