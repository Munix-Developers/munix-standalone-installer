package installer

import (
	"github.com/dchest/uniuri"
	"log"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
	"os"
)

type MountStep struct{}

// Creates a filesystem in each of the partitions.
func (p MountStep) Run(config parser.InstallConfig, context *context.InstallContext) error {
	log.Printf("starting mount step")
	setInstallRoot(context)

	var err error = nil

	for _, d := range config.Storage.Devices {
		for _, p := range d.Partitions {
			setInstallMount(&p, context.GetVar("root"))

			log.Printf("creating %s directories", p.InstallMount)
			err = os.MkdirAll(p.InstallMount, 0755)

			if err != nil {
				return err
			}

			log.Printf("mouting %s in %s%s", p.Device, config.Storage.InstallRoot, p.Mount)
			err = mountDevice(p.Device, p.InstallMount)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Sets the install root to a random directory
func setInstallRoot(c *context.InstallContext) {
	installRoot := uniuri.NewLen(4)
	c.SetVar("root", "/"+installRoot)
}

// Sets the install mount of the PartitionConfig to root + p.Mount
func setInstallMount(p *parser.PartitionConfig, root string) {
	p.InstallMount = root + p.Mount
}

// Mounts a device using the command mount
func mountDevice(device string, mount string) error {
	return utils.StdoutCmd("mount", "-v", device, mount).Run()
}
