package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ClientName string
	TargetInstance string
)

func init() {
	rootCmd.AddCommand(oauthClientsCmd)
	oauthClientsCmd.AddCommand(createOauthClientCmd)
	oauthClientsCmd.AddCommand(deleteOauthClientCmd)
	oauthClientsCmd.AddCommand(listOauthClientCmd)

	oauthClientsCmd.PersistentFlags().StringVarP(&BaseUrl, "baseurl", "b", "https://accounts.bosch-iot-suite.com", "Explicitly set the baseurl of the subscription management API (E.g. '/api/v3/subscriptions' is automatically appended)")
	viper.BindPFlag("baseurl", oauthClientsCmd.PersistentFlags().Lookup("baseurl"))

	createOauthClientCmd.Flags().StringVarP(&ClientName, "clientName", "n", "", "A unique name for the new OAuth2 client")
	createOauthClientCmd.MarkFlagRequired("clientName")
	viper.BindPFlag("clientName", createOauthClientCmd.Flags().Lookup("clientName"))

	createOauthClientCmd.Flags().StringVarP(&TargetInstance, "targetInstance", "i", "", "A service instance name")
	createOauthClientCmd.MarkFlagRequired("targetInstance")
	viper.BindPFlag("targetInstance", createOauthClientCmd.Flags().Lookup("targetInstance"))

	deleteOauthClientCmd.Flags().StringVarP(&ClientName, "clientName", "n", "", "The name of the OAuth client to delete")
	deleteOauthClientCmd.MarkFlagRequired("clientName")
	viper.BindPFlag("clientName", deleteOauthClientCmd.Flags().Lookup("clientName"))

}

var oauthClientsCmd = &cobra.Command{
	Use:   "oauthclients",
	Short: "Manage OAuth2 Clients",
	Long:  `Manage OAuth2 clients for client credentials authorizuation`,
}

var createOauthClientCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new OAuth2 client",
	Long:  `Create a new OAuth2 client for a given service instance`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}

var listOauthClientCmd = &cobra.Command{
	Use:   "list",
	Short: "List all OAuth2 clients",
	Long:  `List all OAuth2 clients`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.ListOAuthClients(httpclient)
	},
}

var deleteOauthClientCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an OAuth2 client",
	Long:  `Delete an OAuth2 client`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.DeleteOAuthClient(httpclient,ClientName)
	},
}

