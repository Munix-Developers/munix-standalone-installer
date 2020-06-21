package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// Creates a cmd that is connected to the system stdout and stderr
func StdoutCmd(name string, arg ...string) *exec.Cmd {
	log.Printf("creating %s %s command", name, strings.Join(arg, " "))

	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
