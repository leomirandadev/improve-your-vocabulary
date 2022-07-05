package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Cache    ConfigCache    `mapstructure:"cache"`
	Database ConfigDatabase `mapstructure:"database"`
}

type ConfigCache struct {
	URL        string        `mapstructure:"url"`
	Expiration time.Duration `mapstructure:"expiration"`
}

type ConfigDatabase struct {
	Reader string `mapstructure:"reader"`
	Writer string `mapstructure:"writer"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
