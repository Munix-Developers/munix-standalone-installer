package validator

import (
	"github.com/go-playground/validator/v10"
	"net.matbm/munix/muinstaller/parser"
)

const acceptedVersion = "1.2"

var validate *validator.Validate

func ValidateConfig(config parser.InstallConfig) error {
	err := getValidator().Struct(config)

	return err
}

func getValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		_ = validate.RegisterValidation("installer-version", installerVersion)
		validate.RegisterStructValidation(bootPartition, parser.StorageConfig{})
	}

	return validate
}

// Validates if the desired config is compatible with current installer.
func installerVersion(f1 validator.FieldLevel) bool {
	version := f1.Field().String()

	return version == acceptedVersion
}

// Validates if the language in the config exists for the Operating System.
func archLanguage(f1 validator.FieldLevel) bool {
	return true // TODO: implement
}

func bootPartition(sl validator.StructLevel) {
	storage := sl.Current().Interface().(parser.StorageConfig)

	bootPartitionCount := 0

	for _, d := range storage.Devices {
		for _, p := range d.Partitions {
			if p.Boot {
				bootPartitionCount++
			}
		}
	}

	if bootPartitionCount != 1 {
		sl.ReportError(storage.Devices, "boot", "Boot", "boot", "")
	}
}
