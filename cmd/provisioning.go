package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(provisioningCmd)
	provisioningCmd.AddCommand(provisionCmd)
	provisioningCmd.AddCommand(myselfCmd)
}

var provisioningCmd = &cobra.Command{
	Use:   "provision",
	Short: "Provision devices",
	Long:  `Device provisioning`,
}

var provisionCmd = &cobra.Command{
	Use:   "custom",
	Short: "Manually register a new device",
	Long:  `Manually register a new device with custom parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}

var myselfCmd = &cobra.Command{
	Use:   "myself",
	Short: "Register this device",
	Long:  `Register this device (eg using current MAC address as deviceId)`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}


