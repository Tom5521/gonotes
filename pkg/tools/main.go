package tools

import (
	"fmt"
	"os"
	"os/user"
)

var HomeDir = func() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
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
		fmt.Println(err)
	}
}

func Chdir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println(err)
	}
}
func Getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
