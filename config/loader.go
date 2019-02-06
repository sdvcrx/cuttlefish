package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	"strings"
)

var (
	configType = "toml"
	config *AppConfig
)

type AppConfig struct {
	Port int
	AuthUser string
	AuthPassword string
	ParentProxies []string
}

func GetInstance() *AppConfig {
	return config
}

func LoadFromString(config string) {
	viper.SetConfigType(configType)
	viper.ReadConfig(strings.NewReader(config))
}

func LoadFromFile(fileName string) {
	viper.SetConfigType(configType)
	viper.SetConfigName(fileName)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrap(err, "config.Load"))
	}
}

func Load() {
	LoadFromFile("config")
	config = &AppConfig{
		Port: viper.GetInt("common.port"),
		AuthUser: viper.GetString("common.username"),
		AuthPassword: viper.GetString("common.password"),
	}
}

func init() {
	pflag.Int("port", 8080, "Proxy server port")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}
