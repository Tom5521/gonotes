package runner

import (
	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
)

func Open() {
	i := FindFileIndexByNameOrID(*flags.Open)
	if i == -1 {
		panic("Could not find note ID or name")
	}
	f := files.Files[i]
	OpenFile(f)
}
