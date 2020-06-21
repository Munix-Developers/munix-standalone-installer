package installer

import "net.matbm/munix/muinstaller/parser"

// Runs the installation steps in a pre-defined order.
func Install(c parser.InstallConfig) error {
	var steps = []InstallStep{
		PartitionsStep{},
		FileSystemStep{},
		MountStep{},
		PacstrapStep{},
	}

	for _, step := range steps {
		err := step.Run(c)

		if err != nil {
			return err
		}
	}

	return nil
}
