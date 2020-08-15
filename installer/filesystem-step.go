package installer

import (
	"log"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
)

type FileSystemStep struct{}

func (p FileSystemStep) GetName() string {
	return "filesystem"
}

// Creates a filesystem in each of the partitions.
func (p FileSystemStep) Run(config parser.InstallConfig, context *context.InstallContext) error {
	var err error = nil

	for _, d := range config.Storage.Devices {
		log.Printf("creating filesystem for %s", d.Device)
		for _, p := range d.Partitions {
			device := context.GetDevice(p)
			log.Printf("creating %s filesystem in %s device", p.Type, device)
			err = createFileSystem(device, p.Type)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates a filesystem using the command "mkfs." + device.
func createFileSystem(device string, fileSystemType string) error {
	return utils.StdoutCmd("mkfs."+fileSystemType, device).Run() // TODO if SWAP do nothing
}
