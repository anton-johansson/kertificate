package cmd

import (
	"fmt"
	"pkims/version"

	"github.com/spf13/cobra"
)

var short bool

func init() {
	var command = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of PKIMS",
		Run: func(command *cobra.Command, args []string) {
			info := version.Info()
			if short {
				fmt.Println(info.Version)
			} else {
				fmt.Println(info.Version + " (go version: " + info.GoVersion + ", commit: " + info.Commit + "), built at " + info.BuildDateHumanReadable())
			}
		},
	}

	command.Flags().BoolVarP(&short, "short", "s", false, "Prints application version only")
	rootCommand.AddCommand(command)
}
