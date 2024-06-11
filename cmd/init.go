package cmd

import (
	"os/user"
	"runtime"
	"strings"

	conf "github.com/Tom5521/goconf"
	"github.com/Tom5521/gonotes/internal/files"
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
		initOpen(),
		initDelete(),
		initConfig(),
		initList(),
	)
}

func MakeFileAndOptions(filename string) (files.File, files.Options) {
	return MakeFile(filename), MakeOptions()
}

func MakeFile(name string) (f files.File) {
	f.Name = name
	f.Type = filetype
	f.Temporal = temporal

	return
}

func MakeOptions() (opts files.Options) {
	switch {
	case temporal:
		opts.NotesPath = settings.String(TemporalPathKey)
	default:
		opts.NotesPath = settings.String(NormalPathKey)
	}

	var suffix string
	switch runtime.GOOS {
	case "windows":
		suffix = "\\"
	default:
		suffix = "/"
	}

	if !strings.HasSuffix(opts.NotesPath, suffix) {
		opts.NotesPath += suffix
	}

	opts.Editor = editor
	return
}
