package geocodemapsyandex

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)


type Config struct {
	APP_ENV  string  `mapstructure:"APP_ENV"`
	APIKEY   string  `mapstructure:"APIKEY"`
	URI      string  `mapstructure:"URI"`
	LANGUAGE string  `mapstructure:"LANGUAGE"`
	LL       string
	SPN      string
	FORMAT   string
}

func LoadConfig(path string) (config Config, err error) {
	// Read file path
	viper.AddConfigPath(path)
	// set config file and path
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	// watching changes in app.env
	viper.AutomaticEnv()
	// reading the config file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func AddresToPoint() {
	config, err := LoadConfig(".")
	// handle errors
	if err != nil {
		log.Fatalf("can't load environment app.env: %v", err)
	}

	fmt.Printf(" -----%s----\n", "Reading Environment variables Using Viper package")
	fmt.Printf(" %s = %v \n", "Application_Environment", config.APP_ENV)
	fmt.Printf(" %s = %v \n", "Yandex API key", config.APIKEY)
	fmt.Printf(" %s = %v \n", "URl", config.URI)
	fmt.Printf(" %s = %v \n", "Speech to text server addres", config.LANGUAGE)
	fmt.Printf(" %s = %v \n", "Selected voice", config.FORMAT)
}