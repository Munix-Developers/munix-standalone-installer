package installer

import (
	"log"
	"net.matbm/munix/muinstaller/parser"
	"os/exec"
)

type FileSystemStep struct{}

// Creates a filesystem in each of the partitions.
func (p FileSystemStep) Run(c parser.InstallConfig) error {
	log.Printf("starting filesystem step")
	var err error = nil

	for _, d := range c.Storage.Devices {
		log.Printf("creating filesystem for %s", d.Device)
		for _, p := range d.Partitions {
			log.Printf("creating %s filesystem in %s device", p.Type, p.Device)
			err = createFileSystem(p.Device, p.Type)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Creates a filesystem using the command "mkfs." + device.
func createFileSystem(device string, fileSystemType string) error {
	return exec.Command("mkfs."+fileSystemType, device).Run()
}
