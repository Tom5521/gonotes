//go:build !release
// +build !release

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func init() {
	cmd := &cobra.Command{
		Use: "test",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return doc.GenMarkdownTree(root, "docs")
		},
	}

	root.AddCommand(cmd)
}
