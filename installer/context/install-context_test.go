package context

import (
	"github.com/stretchr/testify/assert"
	"net.matbm/munix/muinstaller/parser"
	"testing"
)

func TestValueRead(t *testing.T) {
	a := assert.New(t)

	context := New()
	context.SetVar("variable", "value")

	a.Equal("value", context.GetVar("variable"))
}

func TestSetDeviceForPartition(t *testing.T) {
	a := assert.New(t)

	context := New()

	partition := parser.PartitionConfig{
		Type:  "ext4",
		Mount: "/mount",
	}

	context.SetDeviceForPartition(partition, "/dev/sda2")
	a.Equal("/dev/sda2", context.GetDevice(partition))
}

func TestSetInstallMountForPartition(t *testing.T) {
	a := assert.New(t)

	context := New()

	partition := parser.PartitionConfig{
		Type:  "ext4",
		Mount: "/mount",
	}

	context.SetInstallMountForPartition(partition, "/rand")
	a.Equal("/rand/mount", context.GetInstallMount(partition))
}
