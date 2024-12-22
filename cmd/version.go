package cmd

import (
	"fmt"

	"github.com/muhfaris/rocket/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version of rocket",
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("Version: %s", config.Version)
		fmt.Println(version)
	},
}
