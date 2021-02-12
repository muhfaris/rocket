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
		viper.GetString("project")
		c := builder.NewCMDBuilder()
		c.Generate()
		return nil
	},
}

func init() {
	cobra.OnInitialize(initconfig)
	rootCmd.PersistentFlags().StringP("new", "n", "Your Project Name", "name for new project")
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("new"))
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
	if err != nil {
		log.Fatalln("Config application:", err)
	}

	log.Println("using config file:", viper.ConfigFileUsed())
}

// Execute is root function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
