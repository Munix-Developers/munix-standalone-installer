package utils

import (
	"os"
	"os/exec"
)

// Creates a cmd that is connected to the system stdout and stderr
func StdoutCmd(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
