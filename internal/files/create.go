package files

import (
	"errors"
	"os"
)

var (
	ErrAlreadyExists = errors.New("The file already exists!")
)

func Create(file File, opts Options) (err error) {
	fullPath := opts.NotesPath + file.Name + file.Type

	_, err = os.Stat(fullPath)
	if err == nil {
		return ErrAlreadyExists
	}

	if _, err := os.Stat(opts.NotesPath); os.IsNotExist(err) {
		err = os.MkdirAll(opts.NotesPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Create(fullPath)
	if err != nil {
		return
	}

	return Open(file, opts)
}
