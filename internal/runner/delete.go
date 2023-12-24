package runner

import (
	"slices"

	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
)

func Delete() {
	i := FindFileIndexByNameOrID(*flags.Delete)
	if i == -1 {
		panic("Could not find note ID or name")
	}
	files.Files = slices.Delete(files.Files, i, i+1)
	files.Write()
}
