package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func initMd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gen-markdown-tree",
		Short: "Generate markdown documentation into the gived directory.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				err = os.MkdirAll(args[0], os.ModePerm)
				if err != nil {
					return err
				}
			}

			return doc.GenMarkdownTree(root, args[0])
		},
	}

	return cmd
}
