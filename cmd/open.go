package cmd

import (
	"github.com/Tom5521/gonotes/internal/files"
	"github.com/spf13/cobra"
)

func initOpen() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "open",
		Short: "Open a existent file",
		Long:  "Open a existent file that is in the normal or temporal storange.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return files.Open(MakeFileAndOptions(args[0]))
		},
	}

	return cmd
}
