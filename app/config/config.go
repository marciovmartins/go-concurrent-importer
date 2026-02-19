package config

import (
	"github.com/spf13/viper"
)

type CLI struct {
	Env string `mapstructure:"ENV"`

	Name            string  `mapstructure:"CLI_NAME"`
	TagVersion      string  `mapstructure:"CLI_TAG_VERSION"`
	NumWorkers 		int 	`mapstructure:"CLI_NUM_WORKERS"`
	BatchSize  		int 	`mapstructure:"CLI_BATCH_SIZE"`
}

type Database struct {
	Driver   string `mapstructure:"DATABASE_DRIVER"`

	Host    string `mapstructure:"DATABASE_HOST"`
	User    string `mapstructure:"DATABASE_USER"`
	Pass    string `mapstructure:"DATABASE_PASS"`
	DB      string `mapstructure:"DATABASE_DB"`
	Port    string `mapstructure:"DATABASE_PORT"`
	SSLmode string `mapstructure:"DATABASE_SSLMODE"`

}



type Config struct {
	CLI      CLI      `mapstructure:",squash"`
	Database Database `mapstructure:",squash"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	env := viper.GetString("ENV")
	switch env {
	case "test":
		viper.SetConfigName(".env.TEST")
	case "dev", "":
		viper.SetConfigName(".env")
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
