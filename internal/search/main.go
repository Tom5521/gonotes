package search

import (
	"errors"
	"strconv"

	"github.com/Tom5521/gonotes/internal/db"
	"github.com/Tom5521/slicelib"
)

var ErrNotFound = errors.New("file not found")

func byID(files slicelib.Slice[db.File], arg string) (f db.File, err error) {
	id, err := strconv.ParseUint(arg, 10, strconv.IntSize)
	if err != nil {
		return f, err
	}

	cp := files.Copy()
	cp.Filter(func(f db.File) bool {
		return uint64(f.ID) == id
	})

	if cp.IsEmpty() {
		return f, ErrNotFound
	}

	return cp.Elem(0), nil
}

func byName(files slicelib.Slice[db.File], name string) (f db.File, err error) {
	files = files.Copy()

	files.Filter(func(f db.File) bool {
		return f.Name == name
	})

	if files.IsEmpty() {
		return f, ErrNotFound
	}

	return files.Elem(0), nil
}

func ByID(id string) (db.File, error) {
	return byID(db.Files, id)
}

func ByName(name string) (db.File, error) {
	return byName(db.Files, name)
}
