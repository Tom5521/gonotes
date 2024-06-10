package cmd

import (
	"github.com/spf13/cobra"
)

func initConfig() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manages the possible program configurations.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
