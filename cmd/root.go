package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	"os"
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	log "github.com/sirupsen/logrus"
)

var (
	conf *iotsuite.Configuration
	Verbose bool
	Debug bool
	Trace bool
	ConfigFile string
	Quiet bool
	NoColor bool
	Env string
	TokenUrl string
	OutputFile string
)

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

var rootCmd = &cobra.Command{
  Use:   "main",
  Short: "Bosch IoT Suite CLI",
  Long: `A command line tool for interacting with Bosch IoT Suite cloud services.
For more details, please visit https://www.bosch-iot-suite.com/`,
}

func init() {
	cobra.OnInitialize(initRoot)
	
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&Trace, "trace", "x", false, "Enable trace output")
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Hide unnecessary output")
	rootCmd.PersistentFlags().BoolVarP(&NoColor, "nocolor", "g", false, "Disable colors in output")
	rootCmd.PersistentFlags().StringVarP(&ConfigFile, "config", "c", "config", "Configuration base filename (file extension '.yml' is automatically added)")
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "***INTERNAL*** Use a different set of endpoints (Empty for production environment or 'dev' or 'int')")
	rootCmd.PersistentFlags().StringVarP(&TokenUrl, "tokenurl", "t", "https://access.bosch-iot-suite.com/token", "Use OAuth2 Token Endpoint")
	rootCmd.PersistentFlags().StringVarP(&OutputFile, "output", "o", "", "Write response out to file")
	
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("nocolor", rootCmd.PersistentFlags().Lookup("nocolor"))
	viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
	viper.BindPFlag("tokenurl", rootCmd.PersistentFlags().Lookup("tokenurl"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	
}

func initRoot() {
	initLogging();
	initConfig();
	configureLogLevels();
}

func initLogging() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
		DisableLevelTruncation: true,
		ForceColors: true})
	log.SetReportCaller(false)
	log.SetOutput(os.Stdout)
}

func configureLogLevels() {
	if (Trace) {
		log.SetLevel(log.TraceLevel)
	} else if (Debug) {
		log.SetLevel(log.DebugLevel)
	} else if (Verbose) {
		log.SetLevel(log.InfoLevel)
	} else if (Quiet) {
		log.SetLevel(log.ErrorLevel)
	} else {
		// Default log level is WARN
		log.SetLevel(log.WarnLevel)
	}
}

func initConfig() {
	iotsuite.ConfigDefaults()
	if (ConfigFile!="") {
		log.Debug("Reading configuration from file: %s.yml", ConfigFile);
		viper.SetConfigName(ConfigFile) // name of config file (without extension)
		conf = iotsuite.ReadConfig()
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
