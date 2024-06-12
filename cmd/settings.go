package cmd

import (
	"runtime"

	conf "github.com/Tom5521/goconf"
	"github.com/Tom5521/gonotes/internal/db"
)

func initSettings() {
	settings, initerr = conf.New("gonotes")
	if initerr != nil {
		root.PrintErrln(initerr)
		return
	}

	// Default settings:

	if settings.String(NormalPathKey) == "" {
		settings.SetString(NormalPathKey, db.HomeDir+"/.gonotes/")
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
}
