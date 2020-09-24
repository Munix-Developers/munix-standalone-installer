package installer

import (
	"log"
	"net.matbm/munix/muinstaller/installer/context"
	"net.matbm/munix/muinstaller/parser"
	"os"
	"os/signal"
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

	installInterruptHandler(config, steps, installContext)

	err := install(config, steps, installContext)
	if err != nil {
		return err
	}

	cleanup(config, steps, installContext)

	return nil
}

func installInterruptHandler(config parser.InstallConfig, steps []InstallStep, installContext *context.InstallContext) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("Interrupt signal caught, cleaning and exiting...")
			cleanup(config, steps, installContext)
		}
	}()
}

func cleanup(config parser.InstallConfig, steps []InstallStep, installContext *context.InstallContext) {
	for i := len(steps) - 1; i >= 0; i -= 1 {
		log.Printf("Cleaning step %d", i)
		steps[i].Cleanup(config, installContext)
	}
}

func install(config parser.InstallConfig, steps []InstallStep, installContext *context.InstallContext) error {
	for _, step := range steps {
		log.Printf("running %s step", step.GetName())
		err := step.Install(config, installContext)

		if err != nil {
			return err
		}
	}
	return nil
}
