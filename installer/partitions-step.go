package installer

import (
	"fmt"
	"log"
	"net.matbm/munix/muinstaller/parser"
	"os/exec"
	"strings"
)

type PartitionsStep struct{}

func (p PartitionsStep) Run(c parser.InstallConfig) error {
	log.Println("starting partition step")
	var err error = nil

	for _, d := range c.Storage.Devices {
		log.Printf("creating gpt label for %s", d.Device)
		err = createGptLabel(d)

		for _, p := range d.Partitions {
			err = createPartition(d, p)
			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// Creates a partition for a device with a custom partition label
func createPartition(d parser.DeviceConfig, p parser.PartitionConfig) error {
	start := formatStart(p.OffsetBytes)
	end := megaBytes(p.OffsetBytes + p.SizeBytes)
	log.Printf("creating partition: start[%s] end[%s] for %s mountpoint", start, end, p.Mount)
	return exec.Command("parted", "-s", "--align", "optimal",
		d.Device,
		"mkpart",
		getPartLabel(p.Mount),
		start,
		end).Run() // TODO: parse why it failed
}

// Generates a partition label based in the mount point prepending 'mx' and replacing '/' for '.'
// / -> mx.
// /home -> mx.home
// /var/www -> mx.var.www
func getPartLabel(mountPoint string) string {
	return "mx" + strings.ReplaceAll(mountPoint, "/", ".")
}

// Returns 0% when size is 0. Size + "M" if larger.
// Required since parted have a hard time figuring the device align when first sectors of drive are consumed.
func formatStart(size uint64) string {
	if (size) == 0 {
		return "0%"
	} else {
		return megaBytes(size)
	}
}

// Returns size + "M"
func megaBytes(size uint64) string {
	return fmt.Sprintf("%dM", size)
}

// Creates a GPT label for a device
func createGptLabel(d parser.DeviceConfig) error {
	return exec.Command("parted", "-s", d.Device, "mklabel", "gpt").Run() // TODO: parse why it failed
}
