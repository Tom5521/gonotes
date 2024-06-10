package new

import (
	"errors"
	"io"
	"os"
)

var (
	ErrAlreadyExists = errors.New("The file already exists!")
)

func NormalFile(name string, path string) (err error) {
	if _, err := os.Stat(path + name); os.IsExist(err) {
		return ErrAlreadyExists
	}

	file, err := os.Create(path + name)
	io.WriteString(file, "")

	return
}
