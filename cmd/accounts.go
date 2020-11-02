package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	RegistrationEMail string
)

func init() {
	rootCmd.AddCommand(accountsCmd)
	accountsCmd.AddCommand(createAccountCmd)
	accountsCmd.AddCommand(terminateAccountCmd)

	createAccountCmd.Flags().StringVarP(&RegistrationEMail, "email", "m", "", "Your E-Mail address to register")
	createAccountCmd.MarkFlagRequired("email")
	viper.BindPFlag("email", createAccountCmd.Flags().Lookup("email"))

}

var accountsCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage your account",
	Long:  `Manage your personal account and your organization details`,
}

var createAccountCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  `Register a new Bosch IoT Suite account`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}

var terminateAccountCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate an account",
	Long:  `Terminate and close your personal account and your organization`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}


