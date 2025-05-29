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
	cmdproject.NewCMD.Flags().StringP("package", "", "", "package project e.g github.com/muhfaris/myproject")
	viper.BindPFlag("package", cmdproject.NewCMD.Flags().Lookup("package"))

	cmdproject.NewCMD.Flags().StringP("project", "", "", "project name e.g myproject")
	viper.BindPFlag("project", cmdproject.NewCMD.Flags().Lookup("project"))

	// add openapi flag project name
	cmdproject.NewCMD.Flags().StringP("openapi", "", "", "path openapi file")
	viper.BindPFlag("openapi", cmdproject.NewCMD.Flags().Lookup("openapi"))

	// architecture e.g hexagonal
	cmdproject.NewCMD.Flags().StringP("arch", "", "hexagonal", "architecture layout e.g hexagonal, cleancode")
	viper.BindPFlag("arch", cmdproject.NewCMD.Flags().Lookup("arch"))

	// redis
	cmdproject.NewCMD.Flags().StringP("cache", "", "", "cache connection string e.g redis, memory")
	viper.BindPFlag("cache", cmdproject.NewCMD.Flags().Lookup("cache"))

	// db
	cmdproject.NewCMD.Flags().StringP("database", "", "", "db connection string e.g sqlite, mysql, postgres, mongodb")
	viper.BindPFlag("database", cmdproject.NewCMD.Flags().Lookup("database"))

	// docker
	cmdproject.NewCMD.Flags().BoolP("docker", "", false, "generate dockerfile")
	viper.BindPFlag("docker", cmdproject.NewCMD.Flags().Lookup("docker"))

	// config rocket.yaml
	cmdproject.NewCMD.Flags().StringP("config", "c", "", "config file(default is $HOME/config.yaml)")
	viper.BindPFlag("config", cmdproject.NewCMD.Flags().Lookup("config"))
}

// Execute is root function
func Execute() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(cmdproject.NewCMD)
	rootCmd.AddCommand(cmdproject.AddCMD)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
