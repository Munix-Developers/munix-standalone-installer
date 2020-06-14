package installer

import (
	"log"
	"net.matbm/munix/muinstaller/parser"
	"os/exec"
)

type PartitionsStep struct{}

func (p PartitionsStep) Run(c parser.InstallConfig) error {
	log.Println("starting Partition step")
	for _, d := range c.Storage.Devices {
		log.Printf("creating gpt label for %s", d.Device)
		err := exec.Command("parted", "-s", d.Device, "mklabel", "gpt").Run() // TODO: parse why it failed

		if err != nil {
			return err
		}
	}

	return nil
}
