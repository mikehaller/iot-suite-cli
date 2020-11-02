package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	BaseUrl string
	Product string
	InstanceName string
)

func init() {
	rootCmd.AddCommand(subscriptionsCmd)
	subscriptionsCmd.AddCommand(listProductsCmd)
	subscriptionsCmd.AddCommand(listSubscriptionsCmd)
	subscriptionsCmd.AddCommand(createSubscriptionCmd)
	subscriptionsCmd.AddCommand(cancelSubscriptionCmd)
//	subscriptionsCmd.AddCommand(terminateSubscriptionCmd)
//	subscriptionsCmd.AddCommand(renameInstanceCmd)
//	subscriptionsCmd.AddCommand(changePlanCmd)
//	subscriptionsCmd.AddCommand(showBindingCredentialsCmd)
//	subscriptionsCmd.AddCommand(downloadBindingCredentialsCmd)
//	subscriptionsCmd.AddCommand(showServiceStatusHealthCmd)

//	statusCmd.Flags().StringVarP(&Region, "region", "r", "all", "The region of the endpoints (EU-1, EU-2 etc.)")
//	viper.BindPFlag("region", statusCmd.Flags().Lookup("region"))

	subscriptionsCmd.PersistentFlags().StringVarP(&BaseUrl, "baseurl", "b", "https://accounts.bosch-iot-suite.com", "Explicitly set the baseurl of the subscription management API (E.g. '/api/v3/subscriptions' is automatically appended)")
	viper.BindPFlag("baseurl", subscriptionsCmd.PersistentFlags().Lookup("baseurl"))

	createSubscriptionCmd.Flags().StringVarP(&Product, "product", "p", "device-management", "The unique product id or name to provision")
	createSubscriptionCmd.MarkFlagRequired("product")
	viper.BindPFlag("product", createSubscriptionCmd.Flags().Lookup("product"))
	
	createSubscriptionCmd.Flags().StringVarP(&InstanceName, "instanceName", "n", "", "A unique name for the new service instance")
	createSubscriptionCmd.MarkFlagRequired("instanceName")
	viper.BindPFlag("instanceName", createSubscriptionCmd.Flags().Lookup("instanceName"))

	cancelSubscriptionCmd.Flags().StringVarP(&InstanceName, "instanceName", "n", "", "A unique name for the new service instance")
	cancelSubscriptionCmd.MarkFlagRequired("instanceName")
	viper.BindPFlag("instanceName", cancelSubscriptionCmd.Flags().Lookup("instanceName"))

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

var listProductsCmd = &cobra.Command{
	Use:   "products",
	Short: "List available products",
	Long:  `List all products which can be booked in your organisation`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.ProductsList(httpclient)
	},
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new subscription",
	Long:  `Create a new subscription and automatically provision the service instance`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.NewSubscription(httpclient,Product,InstanceName)
	},
}

var cancelSubscriptionCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a subscription",
	Long:  `Cancel the subscription and terminate the service instance`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.CancelSubscription(httpclient,InstanceName)
	},
}

