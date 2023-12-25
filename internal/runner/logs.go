package runner

import (
	"fmt"

	"github.com/Tom5521/GoNotes/internal/files"
	"github.com/gookit/color"
)

func PrintLogs() {
	green := color.FgGreen.Render
	red := color.FgRed.Render
	//yellow := color.FgYellow.Render
	for _, f := range files.Files {
		fmt.Print(green("Name:"), f.Name)
		fmt.Print(red(" Path:"), f.Path)
		fmt.Print(green(" Type:"), f.Type)
		fmt.Println(red(" ID:"), f.ID)
	}
}