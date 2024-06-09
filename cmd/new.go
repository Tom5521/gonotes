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
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	flags := cmd.Flags()

	flags.BoolVar(&overwrite, "overwrite", false, "Overwrites when creating a file if one already exists.")
	flags.StringVar(&editor, "editor", "nano", "Specifies the editor to use.")

	return cmd
}
