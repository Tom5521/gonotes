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
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			options := files.Options{
				Name:      args[0],
				Overwrite: overwrite,
				Type:      filetype,
				Editor:    editor,
			}
			switch {
			case temporal:
				options.Path = settings.String(TemporalPathKey)
			default:
				options.Path = settings.String(NormalPathKey)
			}

			if options.Path[len(options.Path)-1] != "/"[0] {
				options.Path += "/"
			}

			return files.Create(options)
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&overwrite, "overwrite", false,
		"Overwrites when creating a file if one already exists.",
	)

	return cmd
}
