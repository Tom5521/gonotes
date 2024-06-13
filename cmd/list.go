package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/spf13/cobra"
)

func initList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all files detected in normal and temporal storage.",
		Long: `Lists all files detected in normal and temporal storage.
The command also has the following aliases:
"ls", "l"`,
		Args:    cobra.NoArgs,
		Aliases: []string{"ls", "l"},
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, file := range db.Files {
				if file.Temporal && !options.Temporal {
					continue
				}
				fmt.Println("----")
				fmt.Print(file)
			}
			return
		},
	}

	return cmd
}
