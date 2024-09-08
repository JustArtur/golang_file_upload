package cmd

import (
	"client/services/file"
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "Get list of all files",
	Long:  `Get list of all files`,
	Run: func(cmd *cobra.Command, args []string) {
		file.Index()
	},
}

func init() {
	rootCmd.AddCommand(indexCmd)
}
