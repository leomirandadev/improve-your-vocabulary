package envs

import (
	"github.com/spf13/viper"
)

func Load(path string, configPointer any) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(configPointer)
}
