package db

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Tom5521/gonotes/internal/options"
)

func NewID() uint {
	if len(Files) == 0 {
		return 1
	}
	return Files[len(Files)-1].ID + 1
}

var (
	ErrAlreadyExists = errors.New("The file already exists!")
)

func makefile(fullPath string, overwrite bool) (err error) {
	fdir := filepath.Dir(fullPath)

	_, err = os.Stat(fullPath)
	if err == nil && !overwrite {
		return ErrAlreadyExists
	}

	if _, err := os.Stat(fdir); os.IsNotExist(err) {
		err = os.MkdirAll(fdir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err = os.Create(fullPath)
	if err != nil {
		return
	}
	return
}

func Create(name string, overwrite bool) (f File, err error) {
	var path string
	if options.Temporal {
		path = options.TemporalNotesPath
	} else {
		path = options.NotesPath
	}

	fullPath := path + name + options.Filetype
	err = makefile(fullPath, overwrite)
	if err != nil {
		return f, err
	}

	f.Name = name
	f.Type = options.Filetype
	f.Temporal = options.Temporal
	f.ID = NewID()
	f.Path = fullPath

	return
}
