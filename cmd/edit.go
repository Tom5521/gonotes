package cmd

import (
	"github.com/spf13/cobra"
)

func initEdit() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit a existent file",
		Long:  "Edit a existent file that is in the normal or temporal storange.",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return
		},
	}

	return cmd
}
