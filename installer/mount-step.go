package installer

import (
	"fmt"
	"log"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
	"os"
	"time"
)

type MountStep struct{}

func (p MountStep) GetName() string {
	return "mount"
}

// Mounts each partition in ROOT
func (p MountStep) Install(config parser.InstallConfig, context *context.InstallContext) error {
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

func (p MountStep) Cleanup(config parser.InstallConfig, context *context.InstallContext) {
	_ = utils.StdoutCmd("umount", "-v", "-R", context.GetVar("root")).Run()
}

// Sets the install root to a random directory
func setInstallRoot(c *context.InstallContext) {
	c.SetVar("root", getTimestampRoot())
}

// Returns /muinstaller-${epoch-time}
func getTimestampRoot() string {
	return fmt.Sprintf("/muinstaller-%d", time.Now().Unix())
}

// Sets the install mount of the PartitionConfig to root + p.Mount
func setInstallMount(context *context.InstallContext, p parser.PartitionConfig, root string) {
	context.SetInstallMountForPartition(p, root)
}

// Mounts a device using the command mount
func mountDevice(device string, mount string) error {
	return utils.StdoutCmd("mount", "-v", device, mount).Run()
}
