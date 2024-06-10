package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	root = cobra.Command{
		Use:   "gonotes",
		Short: "A note manager for the terminal",
		Long:  "A CLI that allows you to manipulate and manage notes from your terminal using your favorite editor.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			err = cmd.ParseFlags(args)
			if err != nil {
				return
			}

			return
		},
	}
)

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
