package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	_ = viper.ReadInConfig()
	for _, key := range viper.AllKeys() {
		_ = viper.BindEnv(key, strings.ToUpper(strings.Replace(key, ".", "_", -1)))
	}
	viper.AutomaticEnv()
}

func Get(name string) interface{} {
	return viper.Get(name)
}

func GetBool(name string) bool {
	return viper.GetBool(name)
}

func GetInt(name string) int {
	return viper.GetInt(name)
}

func GetInt32(name string) int32 {
	return viper.GetInt32(name)
}

func GetInt64(name string) int64 {
	return viper.GetInt64(name)
}

func GetUint(name string) uint {
	return viper.GetUint(name)
}

func GetUint32(name string) uint32 {
	return viper.GetUint32(name)
}

func GetUint64(name string) uint64 {
	return viper.GetUint64(name)
}

func GetIntSlice(name string) []int {
	return viper.GetIntSlice(name)
}

func GetFloat(name string) float64 {
	return viper.GetFloat64(name)
}

func GetTime(name string) time.Time {
	return viper.GetTime(name)
}

func GetDuration(name string) time.Duration {
	return viper.GetDuration(name)
}

func GetStr(name string) string {
	return viper.GetString(name)
}

func GetStrSlice(name string) []string {
	return viper.GetStringSlice(name)
}

func GetStrMap(name string) map[string]interface{} {
	return viper.GetStringMap(name)
}

func GetStrMapStr(name string) map[string]string {
	return viper.GetStringMapString(name)
}

func GetStrMapStrSlice(name string) map[string][]string {
	return viper.GetStringMapStringSlice(name)
}
