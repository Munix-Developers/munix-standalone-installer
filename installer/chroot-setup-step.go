package installer

import (
	"fmt"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
)

type ChrootSetupStep struct{}

func (p ChrootSetupStep) GetName() string {
	return "chroot"
}

// Creates mount points crucial to a chroot.
// Based in the arch-chroot shell script.
// TODO: add resolv.conf setup
func (p ChrootSetupStep) Run(config parser.InstallConfig, ic *context.InstallContext) error {
	var err error = nil

	root := ic.GetVar("root")

	err = utils.StdoutCmd("mount", "proc", fmt.Sprintf("%s/proc", root), "-v", "-t", "proc", "-o", "nosuid,noexec,nodev").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "sys", fmt.Sprintf("%s/sys", root), "-v", "-t", "sysfs", "-o", "nosuid,noexec,nodev,ro").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "udev", fmt.Sprintf("%s/dev", root), "-v", "-t", "devtmpfs", "-o", "mode=0755,nosuid").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "devpts", fmt.Sprintf("%s/dev/pts", root), "-v", "-t", "devpts", "-o", "mode=1777,nosuid,nodev").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "shm", fmt.Sprintf("%s/dev/shm", root), "-v", "-t", "tmpfs", "-o", "mode=1777,strictatime,nodev,nosuid").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "/run", fmt.Sprintf("%s/run", root), "-v", "--bind").Run()
	if err != nil {
		return err
	}

	err = utils.StdoutCmd("mount", "tmp", fmt.Sprintf("%s/tmp", root), "-v", "-t", "tmpfs", "-o", "mode=1777,strictatime,nodev,nosuid").Run()
	if err != nil {
		return err
	}

	return nil
}
