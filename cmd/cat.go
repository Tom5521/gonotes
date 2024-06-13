package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/search"
	"github.com/spf13/cobra"
)

func initCat() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cat",
		Short: "Performs golang's equivalent of \"cat\" from unix shells to the file.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, v := range args {
				var (
					file    db.File
					content string
				)
				file, err = search.Deep(v)
				if err != nil {
					return
				}
				content, err = file.Content()
				if err != nil {
					return
				}
				fmt.Println(content)
			}
			return
		},
	}

	return cmd
}
