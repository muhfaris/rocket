package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of rocket",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: v1.0.0-rc.13")
	},
}
