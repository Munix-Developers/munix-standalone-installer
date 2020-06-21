package installer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
)

type PacstrapStep struct{}

// Runs pacstrap with the desired pacman mirror.
func (p PacstrapStep) Run(config parser.InstallConfig, ic *context.InstallContext) error {
	log.Printf("starting pacstrap step")

	var err error = nil

	err = setupPacmanMirror(config.Pacman.Mirror)

	if err != nil {
		return err
	}

	err = utils.StdoutCmd("pacstrap", ic.GetVar("root"), "base").Run()

	if err != nil {
		return err
	}

	return nil
}

func setupPacmanMirror(mirror string) error {
	return ioutil.WriteFile("/etc/pacman.d/mirrorlist\n", []byte(fmt.Sprintf("Server = %s", mirror)),
		0644)
}
