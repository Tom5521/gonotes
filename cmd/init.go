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

	if settings.String("normal-path") == "" {
		usr, err := user.Current()
		if err != nil {
			root.PrintErr(err)
		}

		settings.SetString("normal-path", usr.HomeDir+"/.GoNotes/")
	}

	if settings.String("temporal-path") == "" {
		var path string
		switch runtime.GOOS {
		case "windows":
			path = "C:\\Temp\\"
		default:
			path = "/tmp/"
		}
		settings.SetString("temporal-path", path)
	}

	root.SetErrPrefix(color.Red.Render("ERROR:"))

	flags := root.PersistentFlags()
	flags.StringVar(&filetype, "type", settings.String("default-type"), "Specifies the file type.")

	flags.BoolVar(&temporal, "tmp", settings.Bool("default-tmp"),
		"Perform the operation on a file that is specifically located in the temporary directory.")
	flags.BoolVar(&normal, "normal", settings.Bool("default-normal"),
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
