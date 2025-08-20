package appconfig

import (
	"github.com/spf13/viper"
)

type Config struct {
	App_port          int64  `env:"APP_PORT"`
	DB_name           string `env:"DB_NAME"`
	DB_host           string `env:"DB_HOST"`
	DB_port           int64  `env:"DB_PORT"`
	DB_user           string `env:"DB_USER"`
	DB_pass           string `env:"DB_PASS"`
	DB_SSL            string `env:"DB_SSL"`
	MetaDB_Host       string `env:"METADB_HOST"`
	MetaDB_Port       int64  `env:"METADB_PORT"`
	MetaDB_Pass       string `env:"METADB_PASS"`
	MetaDB_DB         int    `env:"METADB_DB"`
	Token_secret      string `env:"TOKEN_SECRET"`
	SMS_Service_ID    string `env:"SMS_SERVICE_ID"`
	SMS_Account_Sid   string `env:"SMS_ACCOUNT_SID"`
	SMS_Service_Token string `env:"SMS_SERVICE_TOKEN"`
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
	return nil
}
