package files

import (
	"errors"
	"os"
)

var (
	ErrAlreadyExists = errors.New("The file already exists!")
)

func Create(opts Options) (err error) {
	fullPath := opts.Path + opts.Name + opts.Type

	if _, err := os.Stat(fullPath); os.IsExist(err) {
		return ErrAlreadyExists
	}

	if _, err := os.Stat(opts.Path); os.IsNotExist(err) {
		err = os.MkdirAll(opts.Path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Create(fullPath)
	if err != nil {
		return
	}

	return Open(opts)
}
