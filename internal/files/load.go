package files

import (
	"bytes"
	"encoding/json"
	"io"
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
	frd *os.File

	HomeDir = func() string {
		usr, err := user.Current()
		if err != nil {
			log.Fatalln(err)
		}
		return usr.HomeDir
	}()

	filesPath = func() string {
		if runtime.GOOS == "windows" {
			return HomeDir + windowsPath
		}
		return HomeDir + unixPath
	}()
)

func CloseFileReadWriter() error {
	return frd.Close()
}

func LoadFiles() (err error) {
	if _, err = os.Stat(filesPath); os.IsNotExist(err) {
		frd, err = os.Create(filesPath)
		return
	} else {
		frd, err = os.Open(filesPath)
	}
	if err != nil {
		return
	}

	bytes, err := io.ReadAll(frd)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &Files)

	return
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

	_, err = io.Copy(frd, bytes.NewReader(data))
	return
}
