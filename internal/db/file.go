package db

import (
	"os"
	"os/exec"
	"strconv"

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

func (f File) String() (str string) {
	render := color.Green.Render

	name := render("Name: ") + f.Name
	ftype := render("Type: ") + f.Type
	id := render("ID: ") + strconv.Itoa(int(f.ID))

	str += name + "\n"
	str += ftype + "\n"
	str += id + "\n"

	return
}

func (f File) Open() (err error) {
	cmd := exec.Command(options.Editor, f.Path)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
