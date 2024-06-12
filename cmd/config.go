package cmd

import (
	"github.com/spf13/cobra"
)

func initConfig() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manages the possible program configurations.",
		Run:   WorkInProgress,
	}

	return cmd
}
