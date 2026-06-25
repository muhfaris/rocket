package cmd

import (
	"github.com/muhfaris/rocket/examples/books-api/cmd/bootstrap"
	"github.com/muhfaris/rocket/examples/books-api/config"
	"github.com/muhfaris/rocket/examples/books-api/internal/adapter/inbound/rest/router"
	portregistry "github.com/muhfaris/rocket/examples/books-api/internal/core/port/inbound/registry"
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
