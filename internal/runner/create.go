package runner

import (
	"fmt"

	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
	t "github.com/Tom5521/GoNotes/pkg/tools"
)

func CreateTmpFile() {
	var tmpF files.File
	t.Chdir("/tmp")
	if t.IsNotExist("GoNotes") {
		t.Mkdir("GoNotes")
	}
	t.Chdir("GoNotes")
	tmpF.Tmp = true
	tmpF.Path = fmt.Sprintf("%s/%s.%s", t.Getwd(), *flags.New, *flags.Type)
	tmpF.Name = *flags.New
	tmpF.ID = files.GetNewID()
	tmpF.Type = fmt.Sprintf(".%s", *flags.Type)
	// For some damn reason the t.IsExist does not work so I just reverse the t.IsNotExists.
	if !t.IsNotExist(tmpF.Path) {
		panic(fmt.Sprintf("The <%s> file already exists, use --open to edit it or --del to delete it.", tmpF.Path))
	}
	OpenFile(tmpF)
	if !t.IsNotExist(tmpF.Path) {
		files.Files = append(files.Files, tmpF)
		files.Write()
	}
}

func CreateFile() {
	var newF files.File
	t.Chdir(t.HomeDir)
	if t.IsNotExist(".GoNotes") {
		t.Mkdir(".GoNotes")
	}
	t.Chdir(".GoNotes")
	newF.Path = fmt.Sprintf("%s/%s.%s", t.HomeDir+"/.GoNotes/", *flags.New, *flags.Type)
	newF.Name = *flags.New
	newF.ID = files.GetNewID()
	newF.Type = *flags.Type

	if !t.IsNotExist(newF.Path) {
		panic(fmt.Sprintf("The <%s> file already exists, use --open to edit it or --del to delete it.", newF.Path))
	}
	OpenFile(newF)
	if !t.IsNotExist(newF.Path) {
		files.Files = append(files.Files, newF)
		files.Write()
	}
}
