package cmd

import (
	"github.com/Tom5521/gonotes/internal/files"
	"github.com/spf13/cobra"
)

func initNew() *cobra.Command {
	var (
		overwrite bool
	)

	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new note.",
		Long:  "Create a new note using your favorite command-line text editor.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			file, options := MakeFileAndOptions(args[0])
			file.Path = ""
			options.Overwrite = overwrite

			return files.Create(file, options)
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&overwrite, "overwrite", false,
		"Overwrites when creating a file if one already exists.",
	)

	return cmd
}
