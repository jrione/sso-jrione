package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Listen             string `mapstructure:"listen"`
		Port               string `mapstructure:"port"`
		Timeout            int    `mapstructure:"timeout"`
		Mode               string `mapstructure:"mode"`
		XApiKey            string `mapstructure:"x-api-key"`
		AccessTokenSecret  string `mapstructure:"access_token_secret"`
		AccessTokenExpiry  int    `mapstructure:"access_token_expiry"`
		RefreshTokenSecret string `mapstructure:"refresh_token_secret"`
		RefreshTokenExpiry int    `mapstructure:"refresh_token_expiry"`
		AppName            string `mapstructure:"app_name"`
	} `mapstructure:"server"`
	Database struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Hostname string `mapstructure:"hostname"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
	Jaeger struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"jaeger"`
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
