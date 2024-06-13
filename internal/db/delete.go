package db

import (
	"errors"
	"slices"
)

func Delete(file File) error {
	for i, f := range Files {
		if f.ID == file.ID {
			Files = slices.Delete(Files, i, i+1)
			return nil
		}
	}

	return errors.New("File not found")
}
