package cmd

import (
	"github.com/mikehaller/iot-suite-cli/iotsuite"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(teamsCmd)
	teamsCmd.AddCommand(inviteMemberCmd)
	teamsCmd.AddCommand(removeMemberCmd)
}

var teamsCmd = &cobra.Command{
	Use:   "team",
	Short: "Manage your team",
	Long:  `Manage your team, invite team members and assign roles`,
}

var inviteMemberCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite a team member",
	Long:  `Invite a team member`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}

var removeMemberCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a team member",
	Long:  `Remove a team member from your organization`,
	Run: func(cmd *cobra.Command, args []string) {
		// conf := iotsuite.ReadConfig()
		iotsuite.ReadConfig()
		// httpclient := iotsuite.InitOAuth(conf)
		// iotsuite.NewOAuthClient(httpclient,ClientName,TargetInstance)
	},
}


