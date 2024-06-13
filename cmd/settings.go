package cmd

import (
	"runtime"
	"strings"

	conf "github.com/Tom5521/goconf"
	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/gonotes/internal/options"
)

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
