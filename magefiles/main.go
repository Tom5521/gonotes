package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

var (
	build    Build
	MainPath = "./cmd/gonotes/main.go"
	toClean  = []string{
		"builds",
	}
)

func Clean() error {
	for _, f := range toClean {
		err := sh.Rm(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Build) LinuxArm() error {
	err := sh.RunWithV(
		map[string]string{
			"GOOS":   "linux",
			"GOARCH": "arm64",
		},
		"go", "build", "-v", "-o", "builds/gonotes-linux-arm64", MainPath,
	)
	return err
}

func (Build) LinuxAmd() error {
	err := sh.RunWithV(
		map[string]string{},
		"go", "build", "-v", "-o", "builds/gonotes-linux-amd64", MainPath,
	)
	return err
}

func (Build) DarwinArm() error {
	err := sh.RunWithV(
		map[string]string{
			"GOOS":   "darwin",
			"GOARCH": "arm64",
		},
		"go", "build", "-v", "-o", "builds/gonotes-darwin-arm64", MainPath,
	)
	return err
}

func (Build) DarwinAmd() error {
	err := sh.RunWithV(
		map[string]string{
			"GOOS":   "darwin",
			"GOARCH": "amd64",
		},
		"go", "build", "-v", "-o", "builds/gonotes-darwin-amd64", MainPath,
	)
	return err
}

func Release() error {
	fmt.Println("LINUX AMD 64")
	err := build.LinuxAmd()
	if err != nil {
		return err
	}
	fmt.Println("LINUX ARM 64")
	err = build.LinuxArm()
	if err != nil {
		return err
	}
	fmt.Println("DARWIN AMD 64")
	err = build.DarwinAmd()
	if err != nil {
		return err
	}
	fmt.Println("DARWIN ARM 64")
	err = build.DarwinArm()
	if err != nil {
		return err
	}
	return nil
}
