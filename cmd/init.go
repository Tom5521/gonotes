package cmd

import (
	"fmt"
	"os/user"
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
	settings, initerr = conf.New("gonotes")
	if initerr != nil {
		root.PrintErrln(initerr)
	}

	// Default settings:

	if settings.String(NormalPathKey) == "" {
		usr, err := user.Current()
		if err != nil {
			root.PrintErr(err)
		}

		settings.SetString(NormalPathKey, usr.HomeDir+"/.gonotes/")
	}

	if settings.String(TemporalPathKey) == "" {
		var path string
		switch runtime.GOOS {
		case "windows":
			path = "C:\\Temp\\gonotes\\"
		default:
			path = "/tmp/gonotes/"
		}
		settings.SetString(TemporalPathKey, path)
	}

	if settings.String(DefaultEditorKey) == "" {
		settings.SetString(DefaultEditorKey, "nano")
	}
	if settings.String(DefaultTypeKey) == "" {
		settings.SetString(DefaultTypeKey, ".txt")
	}

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
	)
}

func InitOptions() {
	if options.Temporal {
		options.NotesPath = settings.String(TemporalPathKey)
	} else {
		options.NotesPath = settings.String(NormalPathKey)
	}

	var suffix string
	switch runtime.GOOS {
	case "windows":
		suffix = "\\"
	default:
		suffix = "/"
	}

	if !strings.HasSuffix(options.NotesPath, suffix) {
		options.NotesPath += suffix
	}
	return
}

func WorkInProgress(cmd *cobra.Command, args []string) {
	fmt.Println("Work in progress!")
}
