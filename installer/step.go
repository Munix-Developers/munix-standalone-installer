package installer

import (
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
)

type InstallStep interface {
	Install(config parser.InstallConfig, context *context.InstallContext) error
	Cleanup(config parser.InstallConfig, context *context.InstallContext)
	GetName() string
}
