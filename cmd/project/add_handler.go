package cmdproject

import (
	addcmd "github.com/muhfaris/rocket/cmd/project/add"
	"github.com/spf13/cobra"
)

var AddCMD = &cobra.Command{
	Use:   "add",
	Short: "Add commands to adding some features to your project",
	Long: `Add commands has subcommands to adding some features to your project including: 
- add handler
  This command will add new handler to your project and include service or repository information from attached to the endpoint.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	AddCMD.AddCommand(addcmd.AddHandlerCMD)
}
