package configs

import (
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer/otel_jaeger"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string              `mapstructure:"port"`
	Env      string              `mapstructure:"env"`
	Cache    cache.Options       `mapstructure:"cache"`
	Tracer   otel_jaeger.Options `mapstructure:"tracer"`
	Database ConfigDatabase      `mapstructure:"database"`
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
