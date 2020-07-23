package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorize the client",
	Long:  `Authorizes the OAuth2 client with the client credentials grant and display the access token`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		iotsuite.Authorize(conf)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().String("clientId", "", "The OAuth2 client id")
	viper.BindPFlag("clientId", authCmd.Flags().Lookup("clientId"))

	authCmd.Flags().String("clientSecret", "", "The OAuth client secret")
	viper.BindPFlag("clientSecret", authCmd.Flags().Lookup("clientSecret"))

	authCmd.Flags().String("scope", "", "The scopes to be requested from the Authorization Server")
	viper.BindPFlag("scope", authCmd.Flags().Lookup("scope"))
}
