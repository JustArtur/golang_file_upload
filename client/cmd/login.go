package cmd

import (
	"client/services/auth"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Use this command for login",
	Long:  `Use this command for login`,
	Run: func(cmd *cobra.Command, args []string) {
		auth.Login(email, password)
	},
}

func init() {
	authCmd.AddCommand(loginCmd)
}
