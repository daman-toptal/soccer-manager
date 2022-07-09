package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var envVars []string

func SetupConfig(envs ...string) {
	viper.AutomaticEnv()
	viper.SetDefault("API_ENV", "sandbox")

	viper.SetConfigType("yml")
	viper.SetConfigName(viper.GetString("API_ENV"))
	viper.SetConfigFile("config/" + viper.GetString("API_ENV") + ".yml")
	err := viper.MergeInConfig()
	if err != nil {
		logrus.Panicf("failed to read config file, err = %s", err.Error())
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}

func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func GetUint(key string) uint {
	return viper.GetUint(key)
}

func GetUint32(key string) uint32 {
	return viper.GetUint32(key)
}

func GetUint64(key string) uint64 {
	return viper.GetUint64(key)
}

func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

func GetIntSlice(key string) []int {
	return viper.GetIntSlice(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
