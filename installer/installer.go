package installer

import (
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
)

// Runs the installation steps in a pre-defined order.
func Install(config parser.InstallConfig) error {
	var steps = []InstallStep{
		PartitionsStep{},
		FileSystemStep{},
		MountStep{},
		PacstrapStep{},
		ChrootSetupStep{},
	}

	installContext := context.New()

	for _, step := range steps {
		err := step.Run(config, installContext)

		if err != nil {
			return err
		}
	}

	return nil
}
