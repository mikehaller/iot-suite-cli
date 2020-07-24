package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(thingsCmd)
	thingsCmd.AddCommand(searchCmd)
	thingsCmd.AddCommand(countCmd)
	thingsCmd.AddCommand(solutionCmd)
	thingsCmd.AddCommand(connectionsCmd)

	searchCmd.Flags().String("filter", "", "Filter query")
	viper.BindPFlag("filter", searchCmd.Flags().Lookup("filter"))

	searchCmd.Flags().String("fields", "thingId,attributes,features", "Filter query")
	viper.BindPFlag("fields", searchCmd.Flags().Lookup("fields"))
	
	searchCmd.Flags().String("namespaces", "", "Limit search to list of comma-separated namespaces")
	viper.BindPFlag("namespaces", searchCmd.Flags().Lookup("namespaces"))
	
	
	countCmd.Flags().String("filter", "", "Filter query")
	viper.BindPFlag("filter", countCmd.Flags().Lookup("filter"))

	countCmd.Flags().String("fields", "thingId,attributes,features", "Filter query")
	viper.BindPFlag("fields", countCmd.Flags().Lookup("fields"))
	
	countCmd.Flags().String("namespaces", "", "Limit search to list of comma-separated namespaces")
	viper.BindPFlag("namespaces", countCmd.Flags().Lookup("namespaces"))
	
	thingsCmd.PersistentFlags().String("solutionId", "", "The Solution Id aka Service Instance ID")
	viper.BindPFlag("solutionId", thingsCmd.PersistentFlags().Lookup("solutionId"))
}

var thingsCmd = &cobra.Command{
	Use:   "things",
	Short: "Access things",
	Long:  `Access the Things API`,
}

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count number of things",
	Long:  `Count the number of things`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.ThingsCount(httpclient, viper.GetString("filter"),viper.GetString("namespaces"))
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for things",
	Long:  `Search for things with a filter`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.Things(httpclient, viper.GetString("fields"), viper.GetString("filter"),viper.GetString("namespaces"))
	},
}

var solutionCmd = &cobra.Command{
	Use:   "solution",
	Short: "Show the configuration of a Solution",
	Long:  `Retrieves and displays the configuration of a Bosch IoT Things Solution`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.ThingsSolution(conf, httpclient, viper.GetString("solutionId"))
	},
}

var connectionsCmd = &cobra.Command{
	Use:   "connections",
	Short: "Show connections",
	Long:  `Shows a list of all configured connections`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.ThingsConnections(conf, httpclient, viper.GetString("solutionId"))
	},
}
