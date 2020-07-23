package iotsuite

import (
	"fmt"
	"github.com/spf13/viper" // yml config
)

func configDefaults() {
	viper.SetEnvPrefix("BOSCH_IOT_")
	viper.AutomaticEnv()
	viper.SetConfigName("config")                 // name of config file (without extension)
	viper.SetConfigType("yml")                   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/bosch-iot-suite/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.bosch-iot-suite") // call multiple times to add many search paths
	viper.AddConfigPath(".")                      // optionally look for config in the working directory

	viper.SetDefault("clientId", "")
	viper.SetDefault("clientSecret", "")
	viper.SetDefault("scope", "")
}

type Configuration struct {
	// OAuth
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	Scope        string `yaml:"scope"`
}

func ReadConfig() *Configuration {
	configDefaults()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println("Configuration file not found, creating configuration file with current set of arguments in current working directory.")
		viper.SafeWriteConfigAs("./config.yml")
	}
	conf := &Configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}

