package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var c config

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
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

type SentryConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	DSN     string `mapstructure:"dsn"`
}

type config struct {
	App      AppConfig     `mapstructure:"app"`
	StatsD   StatsDConfig  `mapstructure:"statsd"`
	Datadog  DatadogConfig `mapstructure:"datadog"`
	Sentry   SentryConfig  `mapstructure:"sentry"`
	DBConfig DBConfig      `mapstructure:"database"`
}

func Load(cfgName, path string) error {
	viper.SetConfigName(cfgName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	return viper.Unmarshal(&c)
}

func GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.App.Address, c.App.Port)
}

func GetStatsDAddress() string {
	return fmt.Sprintf("%s:%s", c.StatsD.Host, c.StatsD.Port)
}

func GetDatadogConfig() DatadogConfig {
	return c.Datadog
}

func GetSentry() SentryConfig {
	return c.Sentry
}

func GetDBConfig() DBConfig {
	return c.DBConfig
}
