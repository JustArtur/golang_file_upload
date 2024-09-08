package cmd

import (
	"client/services/file"
	"github.com/spf13/cobra"
)

var filePath string

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file to server",
	Long:  `Upload file to server`,
	Run: func(cmd *cobra.Command, args []string) {
		file.Upload(filePath)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringVarP(&filePath, "file", "f", "", "file path")
}
