package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Uri      string `mapstructure:"url"`
	Name     string `mapstructure:"name"`
}
type Email struct {
	EmailKey string `mapstructure:"key"`
}
type Config struct {
	Database Database `mapstructure:"database"`
	Email    Email    `mapstructure:"email"`
	Port     string   `mapstructure:"port"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Domain  string   `mapstructure:"domain"`
}
type Jwt struct {
	JwtKey string `mapstructure:"jwtKey"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return &config, err
	}
	return &config, nil
}