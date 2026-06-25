package addcmd

import (
	"fmt"
	"log"
	"os"

	"github.com/muhfaris/rocket/builder"
	"github.com/muhfaris/rocket/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg           config.Config
	AddHandlerCMD = &cobra.Command{
		Use: "handler",
		PreRun: func(cmd *cobra.Command, args []string) {
			// if a config file is found, read it in.
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal(err)
				return
			}

			err = viper.Unmarshal(&cfg)
			if err != nil {
				log.Fatal(err)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			openapifile, err := cmd.Flags().GetString("openapi")
			if err != nil || openapifile == "" {
				return fmt.Errorf("--openapi flag is required")
			}

			operationID, err := cmd.Flags().GetString("operationid")
			if err != nil || operationID == "" {
				return fmt.Errorf("--operationid flag is required")
			}

			ignoreDataResponse := ""
			if v, _ := cmd.Flags().GetBool("ignore-data-response"); v {
				ignoreDataResponse = "true"
			}

			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("get working directory: %w", err)
			}

			fmt.Println("Adding handler to project in", wd)
			fmt.Printf("  OpenAPI spec: %s\n", openapifile)
			fmt.Printf("  OperationId:  %s\n", operationID)

			return builder.AddHandler(wd, openapifile, operationID, ignoreDataResponse)
		},
	}
)

func init() {
	// add argument openapi file
	AddHandlerCMD.Flags().StringP("openapi", "", "", "path openapi file")
	viper.BindPFlag("openapi", AddHandlerCMD.Flags().Lookup("openapi"))

	// add argument spesific openration ID from openapi
	AddHandlerCMD.Flags().StringP("operationid", "", "", "operation id")
	viper.BindPFlag("operationid", AddHandlerCMD.Flags().Lookup("operationid"))

	// add ignore response data
	AddHandlerCMD.Flags().BoolP("ignore-data-response", "", false, "ignore data field in response")
	viper.BindPFlag("ignore-data-response", AddHandlerCMD.Flags().Lookup("ignore-data-response"))
}
