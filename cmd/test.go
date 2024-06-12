package cmd

import (
	"fmt"

	"github.com/Tom5521/gonotes/internal/options"
	"github.com/spf13/cobra"
)

func initTest() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "test",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			fmt.Println(settings.String(NormalPathKey))
			fmt.Printf("\"%s\"\n", options.NotesPath)
			return
		},
	}

	return cmd
}
