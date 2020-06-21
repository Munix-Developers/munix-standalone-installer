package context

import (
	"net.matbm/munix/muinstaller/parser"
)

type InstallContext struct {
	variables map[string]string
}

func New() *InstallContext {
	return &InstallContext{
		variables: make(map[string]string),
	}
}

func (c InstallContext) SetVar(name string, value string) {
	c.variables[name] = value
}

func (c InstallContext) GetVar(name string) string {
	return c.variables[name]
}

func (c InstallContext) SetDeviceForPartition(partition parser.PartitionConfig, device string) {
	c.variables[partition.Mount+partition.Type] = device
}

func (c InstallContext) GetDevice(partition parser.PartitionConfig) string {
	return c.variables[partition.Mount+partition.Type]
}

func (c InstallContext) SetInstallMountForPartition(partition parser.PartitionConfig, root string) {
	c.variables["mount"+partition.Mount+partition.Type] = root + partition.Mount
}

func (c InstallContext) GetInstallMount(partition parser.PartitionConfig) string {
	return c.variables["mount"+partition.Mount+partition.Type]
}
