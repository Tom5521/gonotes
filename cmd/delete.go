package cmd

import "github.com/spf13/cobra"

func initDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a file by its name.",
		Long:  "Delete a file by its name, an error will occur if there are two files with the same name in the normal or temporary storage, also if there are two files with the same name but different file type. In those cases you must specify with a flag.",
	}

	return cmd
}
