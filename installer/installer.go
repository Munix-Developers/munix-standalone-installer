package installer

import "net.matbm/munix/muinstaller/parser"

func Install(c parser.InstallConfig) error {
	return PartitionsStep{}.Run(c)
}
