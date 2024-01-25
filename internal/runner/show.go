package runner

import (
	"fmt"

	"github.com/gookit/color"
)

func PrintConfig() {
	green := color.FgGreen.Render
	fmt.Println(green("Editor:"), conf.Editor)
}
