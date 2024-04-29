package config

import (
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
	Registered   AccountConfig `mapstructure:"registered"`
	Unregistered AccountConfig `mapstructure:"unregistered"`
}

type AccountConfig struct {
	BalanceLimit             int32 `mapstructure:"balance_limit"`
	CreditCountDailyLimit    int16 `mapstructure:"credit_count_daily_limit"`
	CreditCountMonthlyLimit  int16 `mapstructure:"credit_count_monthly_limit"`
	CreditAmountDailyLimit   int64 `mapstructure:"credit_amount_daily_limit"`
	CreditAmountMonthlyLimit int64 `mapstructure:"credit_amount_monthly_limit"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type config struct {
	App      AppConfig `mapstructure:"app"`
	Log      LogConfig `mapstructure:"log"`
	DBConfig DBConfig  `mapstructure:"database"`
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

func GetLogConfig() LogConfig {
	return c.Log
}

func GetDBConfig() DBConfig {
	return c.DBConfig
}
