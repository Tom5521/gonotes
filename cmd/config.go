package cmd

import (
	"github.com/spf13/cobra"
)

func initConfig() *cobra.Command {
	var reset string

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manages the possible program configurations.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	flags := cmd.Flags()
	flags.StringVar(&reset, "reset", "", "Return the key to its default state.")

	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "Returns all keys to default state.",
	}

	cmd.AddCommand(resetCmd)

	return cmd
}
