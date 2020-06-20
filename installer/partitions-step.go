package installer

import (
	"bytes"
	"fmt"
	"log"
	"net.matbm/munix/muinstaller/parser"
	"net.matbm/munix/muinstaller/utils"
	"sort"
	"strings"
)

type PartitionsStep struct{}

func (p PartitionsStep) Run(c parser.InstallConfig) error {
	log.Println("starting partition step")
	var err error = nil

	for _, d := range c.Storage.Devices {
		log.Printf("creating gpt label for %s", d.Device)
		err = createGptLabel(d)

		if err != nil {
			return err
		}

		for _, p := range d.Partitions {
			err = createPartition(d, p)
			if err != nil {
				return err
			}

		}

		err = discoverPartitionDevices(&d)
		if err != nil {
			return err
		}
	}

	return nil
}

// Creates a partition for a device with a custom partition label
func createPartition(d parser.DeviceConfig, p parser.PartitionConfig) error {
	start := formatStart(p.StartMegaBytes)
	end := megaBytes(p.StartMegaBytes + p.SizeMegaBytes)
	log.Printf("creating partition: device %s\tstart %s\tend %s\t\tfor %s mountpoint", d.Device, start, end, p.Mount)
	return utils.StdoutCmd("parted", "-s", "--align", "optimal",
		d.Device,
		"mkpart",
		getPartLabel(p.Mount),
		start,
		end).Run()
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
	return utils.StdoutCmd("parted", "-s", d.Device, "mklabel", "gpt").Run()
}

// Since parted doesn't returns which is the device that was created, we need to find it manually. This function searches
// for partitions where the label starts with "mx" and exists in the configuration
func discoverPartitionDevices(d *parser.DeviceConfig) error {
	blkidOut, err := runBlkid()
	if err != nil {
		return err
	}

	sort.SliceStable(d.Partitions, func(i, j int) bool {
		return d.Partitions[i].Mount <= d.Partitions[j].Mount
	})

	for _, line := range strings.Split(blkidOut, "\n") {
		line = strings.ReplaceAll(line, ": PARTLABEL=", " ")
		line = strings.ReplaceAll(line, "\"", "")
		data := strings.Split(line, " ")

		err = proccessBlkidLine(d, data)
		if err != nil {
			return err
		}
	}

	return nil
}

// Runs blkid searching for the PARTLABEL key in devices
func runBlkid() (string, error) {
	blkid := utils.StdoutCmd("blkid", "-s", "PARTLABEL")
	out := new(bytes.Buffer)
	blkid.Stdout = out

	err := blkid.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

// Checks whether a blkid output have a installer device and mount point. If yes, configure the DeviceConfigPointer. If not, ignore the device.
func proccessBlkidLine(d *parser.DeviceConfig, data []string) error {
	if len(data) > 1 {
		device := data[0]
		label := data[1]

		if strings.HasPrefix(label, "mx") {
			err2 := setDeviceForPartition(d, label, device)
			if err2 != nil {
				return err2
			}
		} else {
			log.Printf("skiping device %s with %s label since it doesn't seems like a installer partition", device, label)
		}
	}
	return nil
}

// Receives a label by blkid, parses it and set the device for a mountpoint inside the DeviceConfig pointer. Halts if the
// device doesn't match any mountpoint inside DeviceConfig
func setDeviceForPartition(d *parser.DeviceConfig, label string, device string) error {
	mount := strings.ReplaceAll(label[2:], ".", "/")

	partId := sort.Search(len(d.Partitions), func(i int) bool {
		return d.Partitions[i].Mount >= mount
	})

	if (partId) == -1 {
		return fmt.Errorf("device for %s mountpoint from %s didn't matched any mounts in partition list, searched mount: %s", device, d.Device, mount)
	}

	partition := &d.Partitions[partId]
	partition.Device = device

	log.Printf("found %s device for %s mount", partition.Device, partition.Mount)
	return nil
}
