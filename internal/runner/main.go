package runner

import (
	"slices"

	"github.com/Tom5521/CmdRunTools/command"
	"github.com/Tom5521/GoNotes/internal/config"
	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/Tom5521/GoNotes/internal/flags"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	t "github.com/Tom5521/GoNotes/pkg/tools"
)

var conf = &config.MainConf
var args = &flags.Args

func Init() {
	flags.InitParsers()
	files.Load()
	CatchTmp()
	switch {
	case args.New != nil:
		switch args.Temporal {
		case true:
			CreateTmpFile()
		default:
			CreateFile()
		}
	case args.Print != "":
		PrintFile()
	case args.Open != "":
		Open()
	case args.List:
		PrintList()
	case args.Config != nil:
		switch {
		case args.Config.Show:
			PrintConfig()
		case args.Config.Editor != conf.Editor:
			conf.Editor = args.Config.Editor
			conf.Update()
		}
	case args.Delete != "":
		Delete()
	}
}

func CatchTmp() {
	for _, f := range files.Files {
		if f.Tmp {
			if t.IsNotExist(f.Path) {
				i := FindFileIndexByID(f.ID)
				if i == -1 {
					continue
				}
				files.Files = slices.Delete(files.Files, i, i+1)
			}
			files.Files = slices.Delete(files.Files, i, i+1)
		}
	}
	files.Write()
}

func OpenFile(f files.File) {
	cmd := command.InitCmdf("%s %s", conf.Editor, f.Path)
	cmd.CustomStd(true, true, true)
	err := cmd.Run()
	if err != nil {
		msg.FatalError(err)
	}
}
