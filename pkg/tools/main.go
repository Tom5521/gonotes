package tools

import (
	"os"
	"os/user"

	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

var HomeDir = func() string {
	usr, err := user.Current()
	if err != nil {
		msg.FatalError(err)
	}
	return usr.HomeDir
}()

func IsExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsExist(err)
}

func IsNotExist(dir string) bool {
	_, err := os.Stat(dir)
	return os.IsNotExist(err)
}

func Mkdir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil {
		msg.FatalError(err)
	}
}

func Chdir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		msg.FatalError(err)
	}
}
func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		msg.FatalError(err)
	}
	return wd
}
