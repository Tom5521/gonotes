package cmd

import (
	"fmt"
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/sahilm/fuzzy"
	"github.com/spf13/cobra"
)

func validNotes(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	arg := args[len(args)-1]

	var names, ids []string
	for _, file := range db.Files {
		ids = append(ids, strconv.FormatUint(uint64(file.ID), 10))
		names = append(names, file.Name)
	}

	matches := fuzzy.Find(arg, names)
	if matches.Len() == 0 {
		matches = fuzzy.Find(arg, ids)
		if matches.Len() == 0 {
			return nil, cobra.ShellCompDirectiveError
		}
	}

	var found []string
	for _, match := range matches {
		found = append(found, match.Str)
	}

	return found, cobra.ShellCompDirectiveNoSpace
}

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
				if !file.Temporal && options.Temporal {
					continue
				}
				if file.Temporal && options.Normal {
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
