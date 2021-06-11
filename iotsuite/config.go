package iotsuite

import (
	"fmt"
	"github.com/spf13/viper" // yml config
	log "github.com/sirupsen/logrus"
)

func ConfigDefaults() {
	log.Debug("Initializing default configuration values")
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
	viper.SetDefault("orgId", "")
	viper.SetDefault("nocolor", "false")
	viper.SetDefault("baseUrl", "https://accounts.bosch-iot-suite.com")
	viper.SetDefault("tokenUrl", "https://access.bosch-iot-suite.com/token")
}

type Configuration struct {
	// OAuth
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
	Scope        string `yaml:"scope"`
	ThingsApiToken string `yaml:"thingsApiToken"`
	ThingsSolutionId string `yaml:"solutionId"`
	ThingsSolutionPolicyId string `yaml:"solutionPolicyId"`
	NoColor        bool `yaml:"nocolor"`
	BaseUrl        string `yaml:"baseUrl"`
	TokenUrl        string `yaml:"tokenUrl"`
	OrgId        string `yaml:"orgId"`
}

func ReadConfig() *Configuration {
	log.Debug("Reading configuration");
	
	err := viper.ReadInConfig() // Find and read the config file
	
	log.WithFields(log.Fields{"settings":viper.AllSettings()}).Trace("Viper settings")
	
	if err != nil {             // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Configuration file not found, creating configuration file with current set of arguments in current working directory.")
			viper.SafeWriteConfigAs("./config.yml")
	    } else {
	        fmt.Printf("Error reading configuration file, %v", err)
	    }
	}
	
	conf := &Configuration{}
	
	err = viper.Unmarshal(conf)

	log.Debug("Configuration: %s",conf);
	
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}

