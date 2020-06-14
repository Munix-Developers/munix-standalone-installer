package installer

import "net.matbm/munix/muinstaller/parser"

type InstallStep interface {
	Run(c parser.InstallConfig) error
}
