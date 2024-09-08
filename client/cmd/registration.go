package cmd

import (
	"client/services/auth"
	"github.com/spf13/cobra"
)

// registrationCmd represents the registration command
var registrationCmd = &cobra.Command{
	Use:   "registration",
	Short: "Command for registration.",
	Long:  `Command for registration.`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Register(email, password)
	},
}

func init() {
	authCmd.AddCommand(registrationCmd)
}
