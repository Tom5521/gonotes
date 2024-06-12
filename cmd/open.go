package cmd

import (
	"github.com/spf13/cobra"
)

func initOpen() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "open",
		Short: "Open a existent file",
		Long:  "Open a existent file that is in the normal or temporal storange.",
		Run:   WorkInProgress,
	}

	return cmd
}
