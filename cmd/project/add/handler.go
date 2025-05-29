package addcmd

import (
	"fmt"
	"log"

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
			openapifile, valid := viper.Get("openapi").(string)
			if !valid {
				return fmt.Errorf("path openapi (--openapi) file is required")
			}

			if openapifile == "" {
				return fmt.Errorf("path openapi (--openapi) file is required")
			}

			// content, doc, err := libos.LoadOpenapi(openapifile)
			// if err != nil {
			// 	return err
			// }
			//
			// operationid := viper.Get("operationid")
			// if operationid == "" {
			// 	return fmt.Errorf("--operationid is required")
			// }
			//
			// ignoreDataResponse, _ := viper.Get("ignore-data-response").(bool)
			//
			// fmt.Println(openapifile)
			// fmt.Println(operationid)

			return nil
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
