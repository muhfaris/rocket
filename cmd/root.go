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
	rootCmd.PersistentFlags().StringP("new", "n", "rocker-sample", "name for new project")
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("new"))

	rootCmd.PersistentFlags().StringP("database", "d", "postgresql", "database for new project (e.g mysql,  postgresql)")
	viper.BindPFlag("db", rootCmd.PersistentFlags().Lookup("new"))

	rootCmd.PersistentFlags().StringP("cache", "e", "none", "cache for new project (e.g redis)")
	viper.BindPFlag("cache", rootCmd.PersistentFlags().Lookup("new"))

	rootCmd.PersistentFlags().StringP("queue", "q", "none", "queue fro new project (e.g rabbitmq)")
	viper.BindPFlag("queue", rootCmd.PersistentFlags().Lookup("new"))

	rootCmd.PersistentFlags().StringP("config", "c", "toml", "config for new project (e.g toml)")
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("new"))
}

func initconfig() {
	viper.SetConfigType("toml")
	// search config in home directory with name "config" (without extension)
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	//read env
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
