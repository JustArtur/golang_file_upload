package cmd

import (
	"client/services/file"
	"github.com/spf13/cobra"
)

var fileName string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download file from server",
	Long:  `Download file from server`,
	Run: func(cmd *cobra.Command, args []string) {
		file.Download(fileName)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVarP(&fileName, "filename", "n", "", "file name")
}
