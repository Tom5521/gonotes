package config

import (
	"encoding/json"
	"os"

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
		panic(err)
	}
	return f
}

func CreateConfigFile() {
	data, err := json.Marshal(Config{Editor: "nano"})
	if err != nil {
		panic(err)
	}
	if t.IsNotExist(ConfigDir) {
		t.Mkdir(ConfigDir)
	}
	file, err := os.Create(ConfigFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
