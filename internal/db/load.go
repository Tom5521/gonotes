package db

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(filesPath)
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

	err = json.Unmarshal(file, &Files)
	return
}

func CatchTmpFiles() error {
	for i, f := range Files {
		if _, err := os.Stat(f.Path); f.Temporal && os.IsNotExist(err) {
			Files = slices.Delete(Files, i, i+1)
			return CatchTmpFiles()
		}
	}

	return WriteFiles()
}

func WriteFiles() (err error) {
	data, err := json.Marshal(Files)
	if err != nil {
		return
	}
	err = os.WriteFile(filesPath, data, os.ModePerm)
	return
}
