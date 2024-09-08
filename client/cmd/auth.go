package cmd

import (
	"github.com/spf13/cobra"
)

var email string
var password string
var Token string

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Group off commands for registration and login.",
	Long:  `Group off commands for registration and login.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "login email")
	authCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "login password")
	authCmd.MarkFlagRequired("email")
	authCmd.MarkFlagRequired("password")
}
