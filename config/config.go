package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var c config

type AppConfig struct {
	Name     string         `mapstructure:"name"`
	Address  string         `mapstructure:"address"`
	Port     string         `mapstructure:"port"`
	Transfer TransferConfig `mapstructure:"transfer"`
}

type TransferConfig struct {
	Registered   AccConfig `mapstructure:"registered"`
	Unregistered AccConfig `mapstructure:"unregistered"`
}

type AccConfig struct {
	BalanceLimit       int32 `mapstructure:"balance_limit"`
	DailyCountLimit    int16 `mapstructure:"daily_count_limit"`
	MonthlyCountLimit  int16 `mapstructure:"monthly_count_limit"`
	DailyAmountlimit   int32 `mapstructure:"daily_amount_limit"`
	MonthlyAmountlimit int32 `mapstructure:"monthly_amount_limit"`
}

type StatsDConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DatadogConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Env     string `mapstructure:"env"`
	Name    string `mapstructure:"service"`
	Version string `mapstructure:"version"`
}

// type SentryConfig struct {
// 	Enabled bool   `mapstructure:"enabled"`
// 	DSN     string `mapstructure:"dsn"`
// }

type config struct {
	App      AppConfig     `mapstructure:"app"`
	StatsD   StatsDConfig  `mapstructure:"statsd"`
	Datadog  DatadogConfig `mapstructure:"datadog"`
	DBConfig DBConfig      `mapstructure:"database"`
}

func Load(cfgName, path string) error {
	viper.SetConfigName(cfgName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.ReadInConfig()
	return viper.Unmarshal(&c)
}

func GetAppConfig() AppConfig {
	return c.App
}

func GetDatadogConfig() DatadogConfig {
	return c.Datadog
}

// func GetSentry() SentryConfig {
// 	return c.Sentry
// }

func GetDBConfig() DBConfig {
	return c.DBConfig
}

func GetStatsDAddress() string {
	return fmt.Sprintf("%s:%s", c.StatsD.Host, c.StatsD.Port)
}
