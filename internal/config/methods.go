package config

import (
	"encoding/json"
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func (c *Config) Update() {
	data, err := json.Marshal(c)
	if err != nil {
		msg.FatalError(err)
	}
	err = os.WriteFile(ConfigFile, data, os.ModePerm)
	if err != nil {
		msg.FatalError(err)
	}
	MainConf = Read()
}
