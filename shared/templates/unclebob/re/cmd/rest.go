package cmd

import (
	"github.com/muhfaris/rocket-examples/config"
	"github.com/muhfaris/rocket-examples/internal/adapter/inbound/rest/routers"
	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run rest server",
	RunE:  restRunE,
}

func restRunE(cmd *cobra.Command, args []string) error {
	api := routers.Init(config.App.Port)
	err := api.Run()
	if err != nil {
		return err
	}

	return nil
}
