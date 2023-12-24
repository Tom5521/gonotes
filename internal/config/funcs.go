package config

import (
	"encoding/json"
	"os"
)

func (c *Config) Update() {
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(ConfigFile, data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	MainConf = Read()
}
