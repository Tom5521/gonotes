package cmd

import (
	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/search"
	"github.com/spf13/cobra"
)

func initDelete() *cobra.Command {
	var ()
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a file by its name.",
		Long:  "Delete a file by its name, an error will occur if there are two files with the same name in the normal or temporary storage, also if there are two files with the same name but different file type. In those cases you must specify with a flag.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, v := range args {
				var file db.File
				file, err = search.Deep(v)
				if err != nil {
					return
				}

				err = db.Delete(file)
				if err != nil {
					return
				}
			}

			return
		},
	}

	return cmd
}
