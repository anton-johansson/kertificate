package cmd

import (
	"pkims/app"

	"github.com/spf13/cobra"
)

func init() {
	var command = &cobra.Command{
		Use:   "start",
		Short: "Starts an instance of PKIMS",
		Run: func(command *cobra.Command, args []string) {
			app.Run()
		},
	}

	rootCommand.AddCommand(command)
}
