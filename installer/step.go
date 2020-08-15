package installer

import (
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
)

type InstallStep interface {
	Run(config parser.InstallConfig, context *context.InstallContext) error
	GetName() string
}
