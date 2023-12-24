package runner

import (
	"os"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/Tom5521/GoNotes/internal/config"
	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
	t "github.com/Tom5521/GoNotes/pkg/tools"
	flag "github.com/spf13/pflag"
)

var conf = &config.MainConf

func Init() {
	flag.Parse()
	if *flags.New == "" {
		panic("The file name has not been specified")
	}
	files.Load()
	if *flags.Temporal {
		CreateTmpFile()
	} else {
		CreateFile()
	}
}

func CreateTmpFile() {
	var tmpF files.File
	t.Chdir("/tmp")
	if t.IsNotExist("GoNotes") {
		t.Mkdir("GoNotes")
	}
	t.Chdir("GoNotes")
	tmpF.Tmp = true
	tmpF.Path = t.Getwd() + *flags.New + *flags.Type
}

func CreateFile() {
	var newF files.File
	t.Chdir(t.HomeDir)
	if t.IsNotExist(".GoNotes") {
		t.Mkdir(".GoNotes")
	}
	t.Chdir(".GoNotes")
	newF.Path = t.HomeDir + "/.GoNotes/" + *flags.New + *flags.Type
	newF.Name = *flags.New
	newF.ID = files.GetNewID()
	newF.Type = *flags.Type

	if t.IsExist(newF.Path) {
		panic("The <%s> file already exists, use --open to edit it or --del to delete it.")
	}
	f, err := os.Create(*flags.New + *flags.Type)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	OpenFile(newF)
}

func OpenFile(f files.File) {
	cmd := command.InitCmdf("%s %s", conf.Editor, f.Path)
	cmd.CustomStd(true, true, true)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
