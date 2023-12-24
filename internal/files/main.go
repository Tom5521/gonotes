package files

import (
	"encoding/json"
	"os"

	t "github.com/Tom5521/GoNotes/pkg/tools"
)

type File struct {
	ID   uint
	Name string
	Path string
	Type string
	Tmp  bool
}

var Files []File

func Read() []byte {
	t.Chdir(t.HomeDir)
	if t.IsNotExist(".config/GoNotes") {
		t.Mkdir(t.HomeDir + "/.config/GoNotes")
	}
	if t.IsNotExist(".config/GoNotes/files.json") {
		err := os.WriteFile(".config/GoNotes/files.json", []byte(""), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	bytedata, err := os.ReadFile(t.HomeDir + "/.config/GoNotes/files.json")
	if err != nil {
		panic(err)
	}
	return bytedata
}
func Load() {
	err := json.Unmarshal(Read(), &Files)
	if err != nil {
		panic(err)
	}
}

func GetNewID() uint {
	if len(Files) == 0 {
		return 1
	} else {
		return Files[len(Files)-1].ID + 1
	}
}
