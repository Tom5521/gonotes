package search

import (
	"errors"

	"github.com/Tom5521/gonotes/internal/db"
)

var ErrNotFound = errors.New("File not found")

func ByID(id uint) (db.File, error) {
	for _, f := range db.Files {
		if f.ID == id {
			return f, nil
		}
	}
	return db.File{}, ErrNotFound
}

func ByName(name string) (db.File, error) {
	for _, f := range db.Files {
		if f.Name == name {
			return f, nil
		}
	}
	return db.File{}, ErrNotFound
}
