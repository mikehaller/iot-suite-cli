package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Region string
	Sort   string
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows service health status",
	Long:  `Shows an overview of the current cloud service health status, whether they are up and OPERATIONAL`,
	Run: func(cmd *cobra.Command, args []string) {
		iotsuite.ShowServiceStatusHealth(Region, Sort, false)
	},
}

func init() {
	statusCmd.Flags().StringVarP(&Region, "region", "r", "all", "The region of the endpoints (EU-1, EU-2 etc.)")
	viper.BindPFlag("region", statusCmd.Flags().Lookup("region"))

	statusCmd.Flags().StringVarP(&Sort, "sort", "s", "name", "Sort the list by service name")
	viper.BindPFlag("sort", statusCmd.Flags().Lookup("sort"))

	rootCmd.AddCommand(statusCmd)
}
