package files

import (
	"os"
	"os/exec"
)

func Open(f File, opts Options) error {
	fullpath := opts.NotesPath + f.Name + f.Type
	cmd := exec.Command(opts.Editor, fullpath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
