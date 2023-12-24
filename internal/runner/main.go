package runner

import (
	"slices"

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
	if *flags.Help {
		flag.PrintDefaults()
		return
	}
	if *flags.Config {
		if *flags.SetDefault != conf.Editor {
			conf.Editor = *flags.SetDefault
			conf.Update()
		}
		return
	}
	files.Load()
	CatchTmp()
	if *flags.Log {
		PrintLogs()
		return
	}
	if *flags.Delete != "" {
		Delete()
		return
	}
	if *flags.Open != "" {
		Open()
		return
	}

	if *flags.New != "" {
		if *flags.Temporal {
			CreateTmpFile()
		} else {
			CreateFile()
		}
		return
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
		}
	}
	files.Write()
}

func OpenFile(f files.File) {
	cmd := command.InitCmdf("%s %s", conf.Editor, f.Path)
	cmd.CustomStd(true, true, true)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
