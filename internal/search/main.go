package search

import (
	"errors"
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
)

var ErrNotFound = errors.New("file not found")

func byID(files []db.File, arg string) (db.File, error) {
	id, err := strconv.ParseUint(arg, 10, strconv.IntSize)
	if err != nil {
		return db.File{}, err
	}

	for _, f := range files {
		if uint64(f.ID) == id {
			return f, nil
		}
	}
	return db.File{}, ErrNotFound
}

func byName(files []db.File, name string) (db.File, error) {
	for _, f := range files {
		if f.Name == name {
			return f, nil
		}
	}
	return db.File{}, ErrNotFound
}

func ByID(id string) (db.File, error) {
	return byID(db.Files, id)
}

func ByName(name string) (db.File, error) {
	return byName(db.Files, name)
}
