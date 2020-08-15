package installer

import (
	"log"
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

	err := runSteps(config, steps, installContext)
	if err != nil {
		return err
	}

	return nil
}

func runSteps(config parser.InstallConfig, steps []InstallStep, installContext *context.InstallContext) error {
	for _, step := range steps {
		log.Printf("running %s step", step.GetName())
		err := step.Run(config, installContext)

		if err != nil {
			return err
		}
	}
	return nil
}
