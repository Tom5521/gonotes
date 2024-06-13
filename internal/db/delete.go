package db

import (
	"errors"
	"os"
	"slices"
)

func Delete(file File) error {
	for i, f := range Files {
		if f.ID == file.ID {
			Files = slices.Delete(Files, i, i+1)
			return os.Remove(f.Path)
		}
	}

	return errors.New("File not found")
}
