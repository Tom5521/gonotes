package main

import (
	"os"

	"github.com/Tom5521/GoNotes/internal/runner"
	"github.com/Tom5521/GoNotes/pkg/messages"
)

func main() {
	if len(os.Args) == 1 {
		messages.FatalCode = 1
		messages.FatalError("No arguments were provided.")
	}
	runner.Init()
}
