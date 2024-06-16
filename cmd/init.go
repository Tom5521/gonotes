package cmd

import (
	conf "github.com/Tom5521/goconf"
	"github.com/gookit/color"
)

var (
	settings conf.Preferences
	initerr  error
)

func init() {
	initSettings()
	initOptions()
	initFlags()
	root.SetErrPrefix(color.Red.Render("ERROR:"))

	root.AddCommand(
		initLicence(),
		initNew(),
		initOpen(),
		initDelete(),
		initConfig(),
		initList(),
		initSearch(),
		initCat(),
		initPrintSettings(),
		initMd(),
	)
}
