package cmd

import (
	"fmt"

	conf "github.com/Tom5521/goconf"
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
		initTest(),
	)
}

func workInProgress(cmd *cobra.Command, args []string) {
	fmt.Println("Work in progress!")
}
