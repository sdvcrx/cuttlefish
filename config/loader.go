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
	ParentProxies Iterator
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

	proxies := viper.GetStringSlice("proxy.parent_proxies")
	// Add http:// prefix
	for idx, rule := range proxies {
		if strings.HasPrefix(rule, "http") == false {
			proxies[idx] = "http://" + rule
		}
	}

	config = &AppConfig{
		Port: viper.GetInt("common.port"),
		AuthUser: viper.GetString("common.username"),
		AuthPassword: viper.GetString("common.password"),
		ParentProxies: NewProxyIterator(proxies),
	}
}

func init() {
	// Bind pflag
	pflag.IntP("port", "p", 8080, "Proxy server port")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// Register alias
	viper.RegisterAlias("common.port", "port")
}
