package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Filter string
	Option string
	Namespaces string
	Fields string
)

func init() {
	rulesCmd.Flags().StringVar(&Filter,"filter", "", "Filter query")
	viper.BindPFlag("filter", rulesCmd.Flags().Lookup("filter"))
	rulesCmd.Flags().StringVar(&Option,"option", "", "Paging operations")
	viper.BindPFlag("option", rulesCmd.Flags().Lookup("option"))
	
	devicesCmd.Flags().StringVar(&Filter, "filter", "", "Filter query")
	viper.BindPFlag("filter", devicesCmd.Flags().Lookup("filter"))
	devicesCmd.Flags().StringVar(&Option, "option", "", "Paging operations")
	viper.BindPFlag("option", devicesCmd.Flags().Lookup("option"))
	devicesCmd.Flags().StringVar(&Namespaces, "namespaces", "", "Limits the query to the given list of comma-separated namespaces")
	viper.BindPFlag("namespaces", devicesCmd.Flags().Lookup("namespaces"))
	devicesCmd.Flags().StringVar(&Fields, "fields", "id,policyId,attributes,features", "Comma separated list of fields to be included in the returned JSON")
	viper.BindPFlag("fields", devicesCmd.Flags().Lookup("fields"))
	
	tasksCmd.Flags().StringVar(&Filter,"filter", "", "Filter query")
	viper.BindPFlag("filter", tasksCmd.Flags().Lookup("filter"))
	tasksCmd.Flags().StringVar(&Option,"option", "", "Paging operations")
	viper.BindPFlag("option", tasksCmd.Flags().Lookup("option"))
	
	groupsCmd.Flags().StringVar(&Filter,"filter", "", "Filter query")
	viper.BindPFlag("filter", groupsCmd.Flags().Lookup("filter"))
	groupsCmd.Flags().StringVar(&Option, "option", "", "Paging operations")
	viper.BindPFlag("option", groupsCmd.Flags().Lookup("option"))
	
	rootCmd.AddCommand(iotmgrCmd)
	
	iotmgrCmd.AddCommand(rulesCmd)
	iotmgrCmd.AddCommand(devicesCmd)
	iotmgrCmd.AddCommand(tasksCmd)
	iotmgrCmd.AddCommand(groupsCmd)
}

var iotmgrCmd = &cobra.Command{
	Use:   "iotmgr",
	Short: "Access devices",
	Long:  `Access the Bosch IoT Manager API`,
}

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "List rules",
	Long:  `List all rules`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.IotmgrRules(httpclient, conf)
	},
}

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "List devices",
	Long:  `List all devices`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.IotmgrDevices(httpclient, Filter, Option, Namespaces, Fields)
	},
}

var tasksCmd = &cobra.Command{
	Use:   "tasks",
	Short: "List tasks",
	Long:  `List all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.IotmgrTasks(httpclient, conf)
	},
}

var groupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List groups",
	Long:  `List all groups`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.IotmgrGroups(httpclient, conf)
	},
}