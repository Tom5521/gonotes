package db

import (
	"errors"

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

func Create(name string, overwrite bool) (f File, err error) {
	var path string
	if options.Temporal {
		path = options.TemporalNotesPath
	} else {
		path = options.NotesPath
	}

	f = File{
		Name:     name,
		Type:     options.Filetype,
		Temporal: options.Temporal,
		ID:       NewID(),
		Path:     path + name + options.Filetype,
	}
	err = f.create(overwrite)
	if err != nil {
		return
	}

	Files = append(Files, f)

	return
}
