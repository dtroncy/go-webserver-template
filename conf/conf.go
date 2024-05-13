package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Param1 string  `mapstructure:"param1"`
	Param2 string  `mapstructure:"param2"`
	Param3 float64 `mapstructure:"param3"`
}

func LoadConfig() (config Configuration) {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return
}
