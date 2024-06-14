package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func initTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "test",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return doc.GenMarkdownTree(root, "docs")
		},
	}

	return cmd
}
