package config

import (
	"fmt"

	"github.com/Tom5521/GoNotes/pkg/tools"
)

const RelativeConfigDir = "%s/.config/GoNotes/"

var (
	ConfigFilename = "config.json"
	ConfigDir      = fmt.Sprintf(RelativeConfigDir, tools.HomeDir)
	ConfigFile     = ConfigDir + ConfigFilename
	MainConf       = func() Config {
		c := Config{}
		c.Read()
		return c
	}()
)
