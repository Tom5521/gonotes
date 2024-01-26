package runner

import (
	"fmt"
	"strings"

	"github.com/Tom5521/GoNotes/internal/files"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	t "github.com/Tom5521/GoNotes/pkg/tools"
)

func CreateTmpFile() {
	t.Chdir("/tmp")
	if t.IsNotExist("GoNotes") {
		t.Mkdir("GoNotes")
	}
	t.Chdir("GoNotes")
	newType := func() string {
		if strings.HasPrefix(args.New.Type, ".") {
			return args.New.Type
		}
		return "." + args.New.Type
	}()
	tmpF := files.File{
		Tmp:  true,
		Path: fmt.Sprintf("%s/%s", t.Getwd(), args.New.Name+newType),
		Name: args.New.Name,
		ID:   files.GetNewID(),
		Type: "." + args.New.Type,
	}
	// For some damn reason the t.IsExist does not work so I just reverse the t.IsNotExists.
	if !t.IsNotExist(tmpF.Path) {
		msg.FatalErrorf("The <%s> file already exists, use --open to edit it or --del to delete it.", tmpF.Path)
	}
	OpenFile(tmpF)
	if !t.IsNotExist(tmpF.Path) {
		files.Files = append(files.Files, tmpF)
		files.Write()
	}
}

func CreateFile() {
	t.Chdir(t.HomeDir)
	if t.IsNotExist(".GoNotes") {
		t.Mkdir(".GoNotes")
	}
	t.Chdir(".GoNotes")
	newType := func() string {
		if strings.HasPrefix(args.New.Type, ".") {
			return args.New.Type
		}
		return "." + args.New.Type
	}()
	newF := files.File{
		Path: t.HomeDir + "/.GoNotes/" + args.New.Name + newType,
		Name: args.New.Name,
		ID:   files.GetNewID(),
		Type: args.New.Type,
	}
	if !t.IsNotExist(newF.Path) {
		msg.FatalErrorf("The <%s> file already exists, use --open to edit it or --del to delete it.", newF.Path)
	}
	OpenFile(newF)
	if !t.IsNotExist(newF.Path) {
		files.Files = append(files.Files, newF)
		files.Write()
	}
}
