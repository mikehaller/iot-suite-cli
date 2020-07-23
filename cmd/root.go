package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var (
	Verbose bool
	Debug bool
	Quiet bool
	NoColor bool
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
	rootCmd.PersistentFlags().BoolVarP(&NoColor, "no-color", "n", false, "Disable colors in output")
}

