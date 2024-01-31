package config

import (
	"encoding/json"
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	t "github.com/Tom5521/GoNotes/pkg/tools"
)

func (c *Config) Write() {
	data, err := json.Marshal(c)
	if err != nil {
		msg.FatalError(err)
	}
	err = os.WriteFile(ConfigFile, data, os.ModePerm)
	if err != nil {
		msg.FatalError(err)
	}
	c.Read()
}

func (c *Config) Read() {
	if t.IsNotExist(ConfigDir) {
		t.Mkdir(ConfigDir)
	}
	if t.IsNotExist(ConfigFile) {
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
	*c = nConf
}
