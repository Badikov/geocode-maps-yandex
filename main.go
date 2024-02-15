package geocodemapsyandex

import (

	"log"
	"net/http"

	"github.com/spf13/viper"
)


type Config struct {
	APP_ENV  string  `mapstructure:"APP_ENV"`
	APIKEY   string  `mapstructure:"APIKEY"`
	URI      string  `mapstructure:"URI"`
	LANGUAGE string  `mapstructure:"LANGUAGE"`
	RSPN     string  `mapstructure:"RSPN"`
	LL       string  `mapstructure:"LL"`
	SPN      string  `mapstructure:"SPN"`
	FORMAT   string  `mapstructure:"FORMAT"`
	RESULTS   string  `mapstructure:"RESULTS"`
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

	log.Printf(" -----%s----\n", "Reading Environment variables Using Viper package")
	log.Printf(" %s = %v \n", "Application_Environment", config.APP_ENV)
	log.Printf(" %s = %v \n", "Yandex API key", config.APIKEY)
	log.Printf(" %s = %v \n", "URl", config.URI)
	log.Printf(" %s = %v \n", "Lang", config.LANGUAGE)
	log.Printf(" %s = %v \n", "Response format", config.FORMAT)

	// client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, config.URI, nil)
	if err != nil {
		log.Fatalf("Can't create request: %v",err)
	}

// 	https://geocode-maps.yandex.ru/1.x
//   ? apikey=<string>
//   & geocode=<string>
//   & lang=<string>
//   & [kind=<string>]
//   & [rspn=<boolean>]
//   & [ll=<number>, <number>]
//   & [spn=<number>, <number>]
//   & [bbox=<number>,<number>~<number>,<number>]
//   & [results=<integer>]
//   & [skip=<integer>]
//   & [uri=<string>]
	// appending to existing query args
	query := req.URL.Query()
	query.Add("apikey", config.APIKEY)
	query.Add("geocode", "")
	query.Add("lang", config.LANGUAGE)
	query.Add("rspn", config.RSPN)
	query.Add("results", config.RESULTS)
	
	

	// assign encoded query string to http request
	req.URL.RawQuery = query.Encode()
	// lowerCorner "40.19509 43.672949"
	// upperCorner "40.203301 43.678908" область красная поляна
	// lowerCorner "40.256149 43.682481"
	// upperCorner "40.26436 43.688439"  область Эстосадок

	log.Println(req.URL.String())

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatalf("can't get response: %v", err)
	// }
	// defer resp.Body.Close()
	// //hier mey be errors
	// responseBody, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalf("Can't read response: %v",err)
	// }

	// log.Println(resp.Status)
	// log.Println(string(responseBody))

}