package config

import (
	"encoding/json"
	"fmt"
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/Tom5521/GoNotes/pkg/tools"
)

const RelativeConfigDir = "%s/.config/GoNotes/"

var (
	ConfigFilename = "config.json"
	ConfigDir      = fmt.Sprintf(RelativeConfigDir, tools.HomeDir)
	ConfigFile     = ConfigDir + ConfigFilename
	MainConf       = Read()
)

func Read() Config {
	if tools.IsNotExist(ConfigDir) {
		tools.Mkdir(ConfigDir)
	}
	if tools.IsNotExist(ConfigFile) {
		CreateConfigFile()
	}
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		msg.FatalError(err)
	}
	nConf := Config{}
	err = json.Unmarshal(data, &nConf)
	if err != nil {
		msg.FatalError(err)
	}
	return nConf
}
