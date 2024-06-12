package cmd

import (
	"fmt"
	"runtime"
	"strings"

	conf "github.com/Tom5521/goconf"
	"github.com/Tom5521/gonotes/internal/options"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	settings conf.Preferences
	initerr  error
)

func init() {
	initSettings()
	initOptions()

	root.SetErrPrefix(color.Red.Render("ERROR:"))

	flags := root.PersistentFlags()
	flags.StringVar(&options.Filetype, "type", settings.String(DefaultTypeKey), "Specifies the file type.")
	flags.StringVar(&options.Editor, "editor", settings.String(DefaultEditorKey), "Specifies the editor to use.")

	flags.BoolVar(&options.Temporal, "tmp", settings.Bool(DefaultTmpKey),
		"Perform the operation on a file that is specifically located in the temporary directory.")
	flags.BoolVar(&options.Normal, "normal", settings.Bool(DefaultNormalKey),
		"Perform the operation on a file that is specifically located in a normal directory.",
	)
	root.MarkFlagsMutuallyExclusive("tmp", "normal")

	root.AddCommand(
		initLicence(),
		initNew(),
		initOpen(),
		initDelete(),
		initConfig(),
		initList(),
		initTest(),
	)
}

func initOptions() {
	options.TemporalNotesPath = settings.String(TemporalPathKey)
	options.NotesPath = settings.String(NormalPathKey)

	checkSuffix := func(v *string) {
		var suffix string
		switch runtime.GOOS {
		case "windows":
			suffix = "\\"
		default:
			suffix = "/"
		}

		if !strings.HasSuffix(*v, suffix) {
			*v += suffix
		}
	}
	checkSuffix(&options.TemporalNotesPath)
	checkSuffix(&options.NotesPath)
}

func WorkInProgress(cmd *cobra.Command, args []string) {
	fmt.Println("Work in progress!")
}
