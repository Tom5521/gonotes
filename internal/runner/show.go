package runner

import (
	"fmt"
	"os"

	"github.com/Tom5521/GoNotes/internal/files"
	msg "github.com/Tom5521/GoNotes/pkg/messages"
	"github.com/gookit/color"
)

func PrintConfig() {
	green := color.FgGreen.Render
	fmt.Println(green("Editor:"), conf.Editor)
}

func PrintFile() {
	i := Look4All(args.Print)
	if i == -1 {
		msg.FatalError("Could not find note ID or name")
	}
	f := files.Files[i]
	color.Green.Print("Name: ")
	fmt.Println(f.Name)
	color.Red.Print("ID: ")
	fmt.Println(f.ID)
	color.Red.Print("Path: ")
	fmt.Println(f.Path)
	color.Green.Print("Type: ")
	fmt.Println(f.Type)
	bytedata, err := os.ReadFile(f.Path)
	if err != nil {
		msg.FatalError(err)
	}
	color.Yellow.Println("--FILE--")
	fmt.Println(string(bytedata))
	color.Yellow.Println("--END--")
}
