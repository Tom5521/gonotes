package config

import (
	"encoding/json"
	"os"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
	t "github.com/Tom5521/GoNotes/pkg/tools"
)

type Config struct {
	Editor string `json:"Editor"`
}

func GetRawFile() *os.File {
	if t.IsNotExist(ConfigFile) {
		CreateConfigFile()
	}
	f, err := os.Open(ConfigFile)
	if err != nil {
		msg.FatalError(err)
	}
	return f
}

func CreateConfigFile() {
	data, err := json.Marshal(Config{Editor: "nano"})
	if err != nil {
		msg.FatalError(err)
	}
	if t.IsNotExist(ConfigDir) {
		t.Mkdir(ConfigDir)
	}
	file, err := os.Create(ConfigFile)
	if err != nil {
		msg.FatalError(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		msg.FatalError(err)
	}
}
