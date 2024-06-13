package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/search"
	"github.com/spf13/cobra"
)

func initSearch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Look for a note specifying specific patterns.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			file, err := search.Deep(args[0])
			if err != nil {
				return
			}
			fmt.Println(file)
			return
		},
	}

	return cmd
}
