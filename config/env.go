package config

import "github.com/spf13/viper"

type Config struct {
	Server struct {
		Listen  string `json:"listen"`
		Port    string `json:"port"`
		Timeout int    `json:"timeout"`
		Mode    string `json:"mode"`
	} `json:"server"`
	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Hostname string `json:"hostname"`
		Port     string `json:"port"`
		Name     string `json:"name"`
	} `json:"database"`
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
