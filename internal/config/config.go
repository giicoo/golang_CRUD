package config

import "github.com/spf13/viper"

type Config struct {
	Host string `mapstructure:"host"`

	Path_to_sql string `mapstructure:"path_to_sql"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
