package command

import (
	"github.com/rochmanramadhann/fazztrack-vehicle/src/configs/serve"
	database "github.com/rochmanramadhann/fazztrack-vehicle/src/database/gorm"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "simple api with golang",
}

func init() {
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
