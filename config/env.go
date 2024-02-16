package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Listen  string `mapstructure:"listen"`
		Port    string `mapstructure:"port"`
		Timeout int    `mapstructure:"timeout"`
		Mode    string `mapstructure:"mode"`
		XApiKey string `mapstructure:"x-api-key"`
	} `mapstructure:"server"`
	Database struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Hostname string `mapstructure:"hostname"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
}

func NewEnv() *Config {
	Env := Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		panic(err)
	}

	return &Env
}
