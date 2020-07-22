package iotsuite

import (
	"os"
	"fmt"
    "github.com/spf13/viper" // yml config
    "github.com/spf13/pflag" // yml config binding to flags
)


func configDefaults() {
	viper.SetEnvPrefix("BOSCH_IOT_")
	viper.AutomaticEnv()
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/bosch-iot-suite/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.bosch-iot-suite")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	
	pflag.Usage = func() {
        fmt.Println("Usage:",os.Args[0],"<command> <options>")
        fmt.Println()
        fmt.Println("Available commands:")
        fmt.Println("\t status \t Show service status health")
        fmt.Println("\t auth \t Authorize using an OAuth2 Client")
        fmt.Println("\t things \t List things")
        fmt.Println()
        fmt.Println("See '"+os.Args[0]+" <command> -help' to read about specific subcommands and options.")
        fmt.Println()
		pflag.PrintDefaults()
        os.Exit(1)
    }
}



type Configuration struct {
	// General
	Verbose bool `yaml:"verbose"`
	Region string `yaml:"region"`
	Sort string `yaml:"sort"`
	Help bool `yaml:"help"`	
	// OAuth
	ClientId string `yaml:"OAuthClientId"`
	ClientSecret     string `yaml:"OAuthClientSecret"`
	Scope     string `yaml:"Scope"`
	// Bosch IoT Things
	Fields string `yaml:"fields"`	
}



func ReadConfig() *Configuration {
	// Command Line Options - Status
	pflag.String("region", "all", "Show only service status of specified region. Examples: 'EU-1', 'EU-2' etc. For a full list, please visit https://developer.bosch-iot-suite.com/regions-and-endpoints/")
	pflag.String("sort", "name", "Sort by 'name' (default) or by 'status'")
	pflag.Bool("verbose", false, "Enable verbose output")
	pflag.String("clientId", "", "OAuth Client ID for client credentials flow")
	pflag.String("clientSecret", "", "OAuth Client Secret")
	pflag.String("scope", "", "Scope in the form service:<service>:<instanceId>/<role>, see https://accounts.bosch-iot-suite.com/oauth2-clients/")
	pflag.String("fields","thingId,attributes,features","Comma separated list of fields to return in Things Search query")
	pflag.Bool("help",false,"Shows help")

	configDefaults();

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		fmt.Println(Warn("Configuration file not found, creating configuration file with current set of arguments in current working directory."));
		viper.SafeWriteConfigAs("./config.yaml") 
	}
	
	conf := &Configuration{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	
	if len(os.Args) < 2 || conf.Help {
		pflag.Usage()
		os.Exit(1)
	}
	
	return conf
}