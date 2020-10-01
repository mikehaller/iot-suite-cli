package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	BaseUrl string
)

func init() {
	rootCmd.AddCommand(subscriptionsCmd)
	subscriptionsCmd.AddCommand(listSubscriptionsCmd)
//	subscriptionsCmd.AddCommand(createSubscriptionCmd)
//	subscriptionsCmd.AddCommand(terminateSubscriptionCmd)
//	subscriptionsCmd.AddCommand(renameInstanceCmd)
//	subscriptionsCmd.AddCommand(changePlanCmd)
//	subscriptionsCmd.AddCommand(showBindingCredentialsCmd)
//	subscriptionsCmd.AddCommand(downloadBindingCredentialsCmd)
//	subscriptionsCmd.AddCommand(showServiceStatusHealthCmd)

//	statusCmd.Flags().StringVarP(&Region, "region", "r", "all", "The region of the endpoints (EU-1, EU-2 etc.)")
//	viper.BindPFlag("region", statusCmd.Flags().Lookup("region"))


	listSubscriptionsCmd.Flags().StringVarP(&BaseUrl, "baseurl", "b", "https://accounts.bosch-iot-suite.com", "Explicitly set the baseurl of the subscription management API (E.g. '/api/v3/subscriptions' is automatically appended)")
	viper.BindPFlag("baseurl", listSubscriptionsCmd.Flags().Lookup("baseurl"))

}

var subscriptionsCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Manage subscriptions",
	Long:  `Manage subscriptions and service instances, upgrade to higher service plans, rename instances etc.`,
}

var listSubscriptionsCmd = &cobra.Command{
	Use:   "list",
	Short: "List your subscriptions",
	Long:  `List all your service subscriptions`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.SubscriptionsList(httpclient)
	},
}

