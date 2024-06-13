package db

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"runtime"
	"slices"
)

var Files []File

const (
	unixPath    = "/.config/gonotes/files.json"
	windowsPath = "/.gonotes/files.json"
)

var (
	HomeDir   string
	filesPath string
)

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	HomeDir = usr.HomeDir

	if runtime.GOOS == "windows" {
		filesPath = HomeDir + windowsPath
	} else {
		filesPath = HomeDir + windowsPath
	}
}

func LoadFiles() (err error) {
	var file []byte
	if _, err = os.Stat(filesPath); os.IsNotExist(err) {
		_, err = os.Create(filesPath)
		if err != nil {
			return
		}
		return LoadFiles()
	} else {
		file, err = os.ReadFile(filesPath)
		if err != nil {
			return
		}
	}

	return json.Unmarshal(file, &Files)
}

func CatchTmpFiles() {
	for i, f := range Files {
		if _, err := os.Stat(f.Path); f.Temporal && os.IsNotExist(err) {
			Files = slices.Delete(Files, i, i+1)
			CatchTmpFiles()
			return
		}
	}
}

func WriteFiles() (err error) {
	data, err := json.Marshal(Files)
	if err != nil {
		return
	}
	return os.WriteFile(filesPath, data, os.ModePerm)
}
