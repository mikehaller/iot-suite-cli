package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version",
  Long:  `Print the version number of this command line tool`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Bosch IoT Suite CLI -- HEAD")
  },
}