package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./data/config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No configuration file loaded - using defaults")
	}
}

// GetString ...
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool ...
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetFloat64 ...
func GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}

// GetInt ...
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetStringMap ...
func GetStringMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

// GetStringMapString ...
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetStringSlice ...
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetTime ...
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

// IsSet ...
func IsSet(key string) bool {
	return viper.IsSet(key)
}
