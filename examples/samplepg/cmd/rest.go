package cmd

import (
	"github.com/muhfaris/rocket/examples/samplepg/cmd/bootstrap"
	"github.com/muhfaris/rocket/examples/samplepg/config"
	"github.com/muhfaris/rocket/examples/samplepg/internal/adapter/inbound/rest/router"
	portregistry "github.com/muhfaris/rocket/examples/samplepg/internal/core/port/inbound/registry"
	"github.com/spf13/cobra"
)

var (
	restCmd = &cobra.Command{
		Use:     "rest",
		Short:   "Run rest server",
		PreRunE: restPreRunE,
		RunE:    restRunE,
	}

	repo portregistry.Repository
)

func restPreRunE(cmd *cobra.Command, args []string) error {
	repo = bootstrap.InitializeRepository()
	return nil
}

func restRunE(cmd *cobra.Command, args []string) error {
	services := bootstrap.InitializeService(repo)
	api := router.Init(config.App.Port, services)
	err := api.Run()
	if err != nil {
		return err
	}

	return nil
}
