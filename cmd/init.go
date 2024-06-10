package cmd

import (
	"os/user"
	"runtime"

	conf "github.com/Tom5521/goconf"
	"github.com/gookit/color"
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
			path = "C:\\Temp\\"
		default:
			path = "/tmp/"
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
	flags.StringVar(&filetype, "type", settings.String(DefaultTypeKey), "Specifies the file type.")
	flags.StringVar(&editor, "editor", settings.String(DefaultEditorKey), "Specifies the editor to use.")

	flags.BoolVar(&temporal, "tmp", settings.Bool(DefaultTmpKey),
		"Perform the operation on a file that is specifically located in the temporary directory.")
	flags.BoolVar(&normal, "normal", settings.Bool(DefaultNormalKey),
		"Perform the operation on a file that is specifically located in a normal directory.",
	)
	root.MarkFlagsMutuallyExclusive("tmp", "normal")

	root.AddCommand(
		initLicence(),
		initNew(),
		initEdit(),
		initDelete(),
		initConfig(),
	)
}
