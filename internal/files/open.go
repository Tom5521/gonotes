package files

import (
	"os"
	"os/exec"
)

func Open(opts Options) error {
	fullpath := opts.Path + opts.Name + opts.Type
	cmd := exec.Command(opts.Editor, fullpath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
