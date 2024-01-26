package flags

import (
	bindata "github.com/Tom5521/GoNotes/internal/bin-data"
	"github.com/alexflint/go-arg"
)

var Args ArgsStr

func InitParsers() {
	arg.MustParse(&Args)
}

type NewCmd struct {
	Name string `arg:"-n,required" help:"Specifies the file name"`
	Type string `arg:"-t" help:"You specify the file type" default:"txt"`
}

type ConfigCmd struct {
	Editor string `arg:"-e" default:"nano" help:"Configures the editor"`
	Show   bool   `arg:"-P" help:"Prints the configuration values"`
}

type ArgsStr struct {
	New      *NewCmd    `arg:"subcommand:new" help:"Create a new text file"`
	Config   *ConfigCmd `arg:"subcommand:config" help:"Configure some program variables"`
	Open     string     `arg:"-o" help:"Open a file for editing or reading"`
	Delete   string     `arg:"-d" help:"Deletes a file"`
	Temporal bool       `arg:"-t" help:"Specifies whether the operation will be done in the temporary or constant directory."`
	List     bool       `arg:"-l" help:"List all files, whether temporary or not"`
}

func (ArgsStr) Version() string {
	return bindata.Version
}

func (ArgsStr) Description() string {
	return "GoNotes is a program to take notes and save them in a temporary or constant space."
}
