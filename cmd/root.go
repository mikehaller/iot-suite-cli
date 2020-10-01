package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	"os"
)

var (
	Verbose bool
	Debug bool
	Quiet bool
	NoColor bool
	Env string
	TokenUrl string
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
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "debug", "d", false, "Enable debug output")
	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Hide unnecessary output")
	rootCmd.PersistentFlags().BoolVarP(&NoColor, "nocolor", "n", false, "Disable colors in output")
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", "***INTERNAL*** Use a different set of endpoints (Empty for production environment or 'dev' or 'int')")
	rootCmd.PersistentFlags().StringVarP(&TokenUrl, "tokenurl", "t", "https://access.bosch-iot-suite.com/token", "***INTERNAL*** Use OAuth2 Token Endpoint")
	
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("nocolor", rootCmd.PersistentFlags().Lookup("nocolor"))
	viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
	viper.BindPFlag("tokenurl", rootCmd.PersistentFlags().Lookup("tokenurl"))
	
}

