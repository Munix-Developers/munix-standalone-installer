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
			setInstallMount(context, p, context.GetVar("root"))

			installMount := context.GetInstallMount(p)
			log.Printf("creating %s directories", installMount)
			err = os.MkdirAll(installMount, 0755)

			if err != nil {
				return err
			}

			device := context.GetDevice(p)
			log.Printf("mouting %s in %s%s", device, context.GetVar("root"), p.Mount)
			err = mountDevice(device, installMount)

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
func setInstallMount(context *context.InstallContext, p parser.PartitionConfig, root string) {
	context.SetInstallMountForPartition(p, root)
}

// Mounts a device using the command mount
func mountDevice(device string, mount string) error {
	return utils.StdoutCmd("mount", "-v", device, mount).Run()
}
