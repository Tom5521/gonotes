package cmd

import (
	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/search"
	"github.com/spf13/cobra"
)

func initOpen() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "open",
		Short:             "Open a existent file",
		Long:              `Open a existent file that is in the normal or temporal storange.`,
		Args:              cobra.MinimumNArgs(1),
		Aliases:           []string{"edit"},
		ValidArgsFunction: validNotes,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, v := range args {
				var file db.File
				file, err = search.Deep(v)
				if err != nil {
					return
				}
				err = file.Open()
				if err != nil {
					return
				}
			}
			return
		},
	}

	return cmd
}
