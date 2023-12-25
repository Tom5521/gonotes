package runner

import (
	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func Open() {
	i := FindFileIndexByNameOrID(*flags.Open)
	if i == -1 {
		msg.FatalError("Could not find note ID or name")
	}
	f := files.Files[i]
	OpenFile(f)
}
