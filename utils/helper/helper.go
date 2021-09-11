package helper

import (
	"os"
	"strings"

	"github.com/spf13/viper"

	"gobpframe/utils/crypto"
)

func IsDebug() bool {
	return viper.GetString("GIN_MODE") == "debug" || strings.Index(viper.GetString("server.RunMode"), "dev") == 0
}

func PasswordHash(raw string) string {
	salt := "demo.server.password.salt.!~2s*^"
	return crypto.Sha1Sum(crypto.Sha1Sum(raw + viper.GetString("server.PasswordSalt") + salt))
}

func IsStdoutRedirectToFile() bool {
	fileInfo, _ := os.Stdout.Stat()
	return (fileInfo.Mode() & os.ModeCharDevice) == 0
}
