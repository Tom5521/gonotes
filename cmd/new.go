package cmd

import "github.com/spf13/cobra"

func initNew() *cobra.Command {
	var (
		editor    string
		overwrite bool
	)

	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new note.",
		Long:  "Create a new note using your favorite command-line text editor.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			switch {
			case temporal:
			default:
			}

			return
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&overwrite, "overwrite", settings.Bool("auto-overwrite"),
		"Overwrites when creating a file if one already exists.",
	)
	flags.StringVar(&editor, "editor", settings.String("default-editor"), "Specifies the editor to use.")

	return cmd
}
