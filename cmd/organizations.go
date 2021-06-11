package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(organizationsCmd)
	organizationsCmd.AddCommand(listOrgsCmd)

	organizationsCmd.PersistentFlags().StringVarP(&BaseUrl, "baseurl", "b", "https://accounts.bosch-iot-suite.com", "Explicitly set the baseurl of the organizations management API (E.g. '/api/v3/organizations' is automatically appended)")
	viper.BindPFlag("baseurl", organizationsCmd.PersistentFlags().Lookup("baseurl"))
}

var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Manage organizations",
	Long:  `Manage organizations and team members`,
}

var listOrgsCmd = &cobra.Command{
	Use:   "list",
	Short: "List your organizations",
	Long:  `List all organizations where you are a member of and have access to`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := iotsuite.ReadConfig()
		httpclient := iotsuite.InitOAuth(conf)
		iotsuite.OrgList(httpclient)
	},
}

