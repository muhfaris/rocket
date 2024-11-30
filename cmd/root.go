package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/muhfaris/rocket/builder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "rocket",
	RunE: func(cmd *cobra.Command, args []string) error {
		project := viper.GetString("project")
		db := viper.GetString("db")
		cache := viper.GetString("cache")
		queue := viper.GetString("queue")
		config := viper.GetString("config")

		b := builder.NewBuilder(project, db, cache, queue, config)
		b.Generate()

		return nil
	},
}

func init() {
	cobra.OnInitialize(initconfig)

	openapiCmd.Flags().StringP("package", "p", "", "package project e.g github.com/muhfaris/myproject")
	viper.BindPFlag("package", openapiCmd.Flags().Lookup("package"))

	openapiCmd.Flags().StringP("project", "n", "", "project name e.g myproject")
	viper.BindPFlag("project", openapiCmd.Flags().Lookup("project"))

	// add openapi flag project name
	openapiCmd.Flags().StringP("openapi", "o", "", "path openapi file")
	viper.BindPFlag("openapi", openapiCmd.Flags().Lookup("openapi"))

	// architecture e.g hexagonal
	openapiCmd.Flags().StringP("arch", "a", "hexagonal", "architecture layout e.g hexagonal, cleancode")
	viper.BindPFlag("arch", openapiCmd.Flags().Lookup("arch"))

	// redis
	openapiCmd.Flags().StringP("cache", "c", "", "cache connection string e.g redis, memory")
	viper.BindPFlag("cache", openapiCmd.Flags().Lookup("cache"))

	// db
	openapiCmd.Flags().StringP("db", "d", "", "db connection string e.g sqlite, mysql, postgres, mongodb")
	viper.BindPFlag("db", openapiCmd.Flags().Lookup("db"))

	// docker
	openapiCmd.Flags().BoolP("docker", "", false, "generate dockerfile")
	viper.BindPFlag("docker", openapiCmd.Flags().Lookup("docker"))
	rootCmd.AddCommand(openapiCmd)
}

func initconfig() {
	viper.SetConfigType("toml")
	// search config in home directory with name "config" (without extension)
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	// read env
	viper.AutomaticEnv()

	// if a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		log.Println("using config file:", viper.ConfigFileUsed())
	}
}

// Execute is root function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
