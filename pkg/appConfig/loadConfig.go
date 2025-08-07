package appconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App_port     int64  `env:"APP_PORT"`
	DB_name      string `env:"DB_NAME"`
	DB_host      string `env:"DB_HOST"`
	DB_port      int64  `env:"DB_PORT"`
	DB_user      string `env:"DB_USER"`
	DB_pass      string `env:"DB_PASS"`
	DB_SSL       string `env:"DB_SSL"`
	Token_secret string `env:"TOKEN_SECRET"`
}

func LoadConfig(c *Config) error {
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	fmt.Printf("%+v", c)
	return nil
}
