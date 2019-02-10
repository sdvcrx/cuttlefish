package config

import (
	"fmt"
	"github.com/sdvcrx/cuttlefish/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

var (
	appName    = "cuttlefish"
	configType = "toml"
	config     *AppConfig
	logger     = log.Logger
)

type AppConfig struct {
	Host          string
	Port          int
	AuthUser      string
	AuthPassword  string
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
	// => /etc/cuttlefish/config.toml
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
	// => /$HOME/.config/cuttlefish/config.toml
	viper.AddConfigPath(fmt.Sprintf("/$HOME/.config/%s/", appName))
	viper.AddConfigPath(".")

	configFilePath := viper.GetString("config")
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		logger.Warn().Msg("config file not found, using default config and args")
	}
}

func Load() {
	LoadFromFile("config")

	proxies := viper.GetStringSlice("proxy.parent_proxies")
	// Add http:// prefix
	for idx, rule := range proxies {
		if !strings.HasPrefix(rule, "http") {
			proxies[idx] = "http://" + rule
		}
	}

	config = &AppConfig{
		Host:          viper.GetString("common.host"),
		Port:          viper.GetInt("common.port"),
		AuthUser:      viper.GetString("common.username"),
		AuthPassword:  viper.GetString("common.password"),
		ParentProxies: NewProxyIterator(proxies),
	}
}

func init() {
	// Bind pflag
	pflag.StringP("host", "H", "", "Proxy server host")
	pflag.IntP("port", "p", 8080, "Proxy server port")
	pflag.BoolP("version", "V", false, "Show version number and quit")
	pflag.BoolP("verbose", "v", false, "Log more details")
	pflag.StringP("config", "c", "", "Config file path")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	// Register alias
	viper.RegisterAlias("common.host", "host")
	viper.RegisterAlias("common.port", "port")
}
