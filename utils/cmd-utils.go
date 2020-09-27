package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// Creates a cmd that is connected to the system stdout, stdin and stderr
func StdoutCmd(name string, arg ...string) *exec.Cmd {
	log.Printf("creating %s %s command", name, strings.Join(arg, " "))

	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd
}

// Writes the output of a cmd to a file
func StdoutToFile(fileName string, name string, arg ...string) {
	log.Printf("creating %s %s command", name, strings.Join(arg, " "))

	cmd := exec.Command(name, arg...)

	outFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	cmd.Stdout = outFile

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	cmd.Wait()
}

// Creates a cmd that is connected to the system stdout, stdin and stderr, but chrooted to a specific folder
func ChrootedCmd(name string, chrootPath string, arg ...string) *exec.Cmd {
	cmd := StdoutCmd(name, arg...)

	err := syscall.Chroot(chrootPath)
	if err != nil {
		log.Printf("failed to chroot in %s - can't continue", chrootPath)
		panic(err)
	}

	err = syscall.Chdir(chrootPath)
	if err != nil {
		log.Printf("failed to chdir in %s - can't continue", chrootPath)
		panic(err)
	}

	log.Printf("chrooted %s cmd in %s", name, chrootPath)

	return cmd
}
