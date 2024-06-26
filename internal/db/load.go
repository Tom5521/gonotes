package db

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"slices"
)

var Files []File

func TemporalFiles() (files []File) {
	for _, f := range Files {
		if f.Temporal {
			files = append(files, f)
		}
	}
	return
}

func NormalFiles() (files []File) {
	for _, f := range Files {
		if !f.Temporal {
			files = append(files, f)
		}
	}
	return
}

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
	if _, err = os.Stat(filepath.Dir(filesPath)); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(filesPath), os.ModePerm)
		if err != nil {
			return
		}
	}
	if _, err = os.Stat(filesPath); os.IsNotExist(err) {
		err = WriteFiles()
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

func CatchBadFiles() {
	for i, f := range Files {
		if _, err := os.Stat(f.Path); os.IsNotExist(err) {
			Files = slices.Delete(Files, i, i+1)
			CatchBadFiles()
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
