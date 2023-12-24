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

var FilesDir string = t.HomeDir + "/.config/GoNotes/files.json"

var Files []File

func Read() []byte {
	t.Chdir(t.HomeDir)
	if t.IsNotExist(".config/GoNotes") {
		t.Mkdir(".config/GoNotes")
	}
	if t.IsNotExist(FilesDir) {
		bytedata, err := json.Marshal(Files)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(FilesDir, bytedata, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	bytedata, err := os.ReadFile(FilesDir)
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

func Write() {
	bytedata, err := json.Marshal(Files)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(FilesDir, bytedata, os.ModePerm)
	if err != nil {
		panic(err)
	}
	Load()
}

func GetNewID() uint {
	if len(Files) == 0 {
		return 1
	} else {
		return Files[len(Files)-1].ID + 1
	}
}
