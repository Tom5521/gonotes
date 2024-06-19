package db

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/Tom5521/gonotes/internal/options"
	"github.com/gookit/color"
)

type File struct {
	Name     string
	Type     string
	Path     string
	ID       uint
	Temporal bool
}

func (f File) Content() (string, error) {
	data, err := os.ReadFile(f.Path)
	return string(data), err
}

func (f File) String() (str string) {
	render := func(title string, content ...any) {
		r := color.Green.Render
		str += r(title+": ") + fmt.Sprint(content...) + "\n"
	}

	render("Name", f.Name)
	render("Type", f.Type)
	render("ID", f.ID)

	if f.Temporal {
		render("Temporal", f.Temporal)
	}

	return
}

func (f File) create(overwrite bool) (err error) {
	fdir := filepath.Dir(f.Path)

	_, err = os.Stat(f.Path)
	if err == nil && !overwrite {
		return ErrAlreadyExists
	}

	if _, err := os.Stat(fdir); os.IsNotExist(err) {
		err = os.MkdirAll(fdir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	_, err = os.Create(f.Path)
	if err != nil {
		return
	}
	return
}

func (f File) Open() (err error) {
	cmd := exec.Command(options.Editor, f.Path)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (file File) Delete() (err error) {
	err = os.Remove(file.Path)
	if err != nil {
		return
	}
	for i, f := range Files {
		if file.ID == f.ID {
			Files = slices.Delete(Files, i, i+1)
			return
		}
	}

	return fmt.Errorf("File not found")
}
