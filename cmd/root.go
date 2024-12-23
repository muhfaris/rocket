package cmd

import (
	"fmt"
	"os"

	cmdproject "github.com/muhfaris/rocket/cmd/project"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "rocket",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("what to do")
		return nil
	},
}

func init() {
	cmdproject.OpenapiCMD.Flags().StringP("package", "", "", "package project e.g github.com/muhfaris/myproject")
	viper.BindPFlag("package", cmdproject.OpenapiCMD.Flags().Lookup("package"))

	cmdproject.OpenapiCMD.Flags().StringP("project", "", "", "project name e.g myproject")
	viper.BindPFlag("project", cmdproject.OpenapiCMD.Flags().Lookup("project"))

	// add openapi flag project name
	cmdproject.OpenapiCMD.Flags().StringP("openapi", "", "", "path openapi file")
	viper.BindPFlag("openapi", cmdproject.OpenapiCMD.Flags().Lookup("openapi"))

	// architecture e.g hexagonal
	cmdproject.OpenapiCMD.Flags().StringP("arch", "", "hexagonal", "architecture layout e.g hexagonal, cleancode")
	viper.BindPFlag("arch", cmdproject.OpenapiCMD.Flags().Lookup("arch"))

	// redis
	cmdproject.OpenapiCMD.Flags().StringP("cache", "", "", "cache connection string e.g redis, memory")
	viper.BindPFlag("cache", cmdproject.OpenapiCMD.Flags().Lookup("cache"))

	// db
	cmdproject.OpenapiCMD.Flags().StringP("database", "", "", "db connection string e.g sqlite, mysql, postgres, mongodb")
	viper.BindPFlag("database", cmdproject.OpenapiCMD.Flags().Lookup("database"))

	// docker
	cmdproject.OpenapiCMD.Flags().BoolP("docker", "", false, "generate dockerfile")
	viper.BindPFlag("docker", cmdproject.OpenapiCMD.Flags().Lookup("docker"))

	// config rocket.yaml
	cmdproject.OpenapiCMD.Flags().StringP("config", "c", "", "config file(default is $HOME/config.yaml)")
	viper.BindPFlag("config", cmdproject.OpenapiCMD.Flags().Lookup("config"))
}

// Execute is root function
func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cmdproject.OpenapiCMD)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
