package config

import (
	"encoding/json"
)

func (c *Config) Update() {
	data, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	_, err = RawConfigFile.Write(data)
	if err != nil {
		panic(err)
	}
	RawConfigFile = GetRawFile()
	MainConf = Read()
}
