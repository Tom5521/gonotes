package runner

import (
	"slices"

	"github.com/Tom5521/GoNotes/internal/files"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
)

func Delete() {
	i := FindFileIndexByNameOrID(args.Delete)
	if i == -1 {
		msg.PanicError("Could not find note ID or name")
	}
	files.Files = slices.Delete(files.Files, i, i+1)
	files.Write()
}
