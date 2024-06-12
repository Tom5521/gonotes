package cmd

import (
	"fmt"
	"os"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/spf13/cobra"
)

const (
	NormalPathKey    = "normal-path"
	TemporalPathKey  = "temporal-path"
	DefaultTmpKey    = "default-tmp"
	DefaultNormalKey = "default-normal"
	DefaultTypeKey   = "default-type"
	DefaultEditorKey = "default-editor"
)

var (
	root = cobra.Command{
		Use:   "gonotes",
		Short: "A note manager for the terminal",
		Long:  "A CLI that allows you to manipulate and manage notes from your terminal using your favorite editor.",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) (err error) {
			err = cmd.ParseFlags(args)
			if err != nil {
				return
			}
			err = db.LoadFiles()
			if err != nil {
				return
			}
			err = db.CatchTmpFiles()
			if err != nil {
				return
			}
			InitOptions()
			return
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return db.CloseFileReadWriter()
		},
	}
)

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
