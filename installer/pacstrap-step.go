package installer

import (
	"fmt"
	"io/ioutil"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
)

type PacstrapStep struct{}

func (p PacstrapStep) GetName() string {
	return "pacstrap"
}

// Runs pacstrap with the desired pacman mirror.
func (p PacstrapStep) Install(config parser.InstallConfig, ic *context.InstallContext) error {
	var err error = nil

	err = setupPacmanMirror(config.Pacman.Mirror)

	if err != nil {
		return err
	}

	root := ic.GetVar("root")
	err = utils.StdoutCmd("pacstrap", root, "base").Run()

	if err != nil {
		return err
	}

	fstabFile := fmt.Sprintf("%s/etc/fstab", root)
	utils.StdoutToFile(
		fstabFile,
		"genfstab",
		"-U",
		"-p",
		root,
		fmt.Sprintf("%s/etc/fstab", root),
	)

	return nil
}

func (p PacstrapStep) Cleanup(config parser.InstallConfig, context *context.InstallContext) {

}

func setupPacmanMirror(mirror string) error {
	return ioutil.WriteFile("/etc/pacman.d/mirrorlist", []byte(fmt.Sprintf("Server = %s\n", mirror)),
		0644)
}
