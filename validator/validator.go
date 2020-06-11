package validator

import (
	"github.com/go-playground/validator/v10"
	"net.matbm/munix/installer/parser"
)

const accepted_version = "1.0"

var validate *validator.Validate

func ValidateConfig(config parser.InstallConfig) (bool, error) {
	err := getValidator().Struct(config)

	if err != nil {
		return false, err
	}

	return true, nil
}

func getValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		_ = validate.RegisterValidation("arch-language", archLanguage)
		_ = validate.RegisterValidation("installer-version", archLanguage)
	}

	return validate
}

// Validates if the desired config is compatible with current installer.
func installerVersion(f1 validator.FieldLevel) bool {
	version := f1.Field().String()

	return version == accepted_version
}

// Validates if the language in the config exists for the Operating System.
func archLanguage(f1 validator.FieldLevel) bool {
	return true // TODO: implement
}
