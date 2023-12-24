package runner

import (
	"slices"
	"strconv"

	"github.com/Tom5521/GoNotes/internal/files"
)

func FindFileIndexByName(name string) int {
	return slices.IndexFunc(files.Files, func(f files.File) bool {
		return f.Name == name
	})
}

func FindFileIndexByID(id uint) int {
	return slices.IndexFunc(files.Files, func(f files.File) bool {
		return f.ID == id
	})
}

func FindFileIndexByNameOrID(nameOrID string) int {
	i := FindFileIndexByName(nameOrID)
	if i == -1 {
		id, err := strconv.Atoi(nameOrID)
		if err != nil {
			return -1
		}
		if id <= 0 {
			return -1
		}
		i = FindFileIndexByID(uint(id))
	}
	return i
}
