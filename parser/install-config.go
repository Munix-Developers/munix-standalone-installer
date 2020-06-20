package parser

type InstallConfig struct {
	Version      string             `json:"version" validate:"required,installer-version"`
	Keyboard     KeyboardConfig     `json:"keyboard" validate:"required"`
	Localization LocalizationConfig `json:"localization" validate:"required"`
	Computer     ComputerConfig     `json:"computer" validate:"required"`
	Storage      StorageConfig      `json:"storage" validate:"required"`
}

type KeyboardConfig struct {
	Layout string `json:"layout" validate:"required"`
}

type LocalizationConfig struct {
	SystemLanguage string         `json:"system_language" validate:"required"`
	Locales        []LocaleConfig `json:"locales" validate:"required"`
	Timezone       string         `json:"timezone" validate:"required"`
}

type LocaleConfig struct {
	Language string `json:"language" validate:"required,arch-language"`
	Encoding string `json:"encoding" validate:"required"`
}

type ComputerConfig struct {
	Hostname string `json:"hostname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"Password" validate:"required"`
}

type StorageConfig struct {
	InstallRoot string
	Devices     []DeviceConfig `json:"devices" validate:"required"`
}

type DeviceConfig struct {
	Device     string            `json:"device" validate:"required"`
	Partitions []PartitionConfig `json:"partitions" validate:"required"`
}

type PartitionConfig struct {
	Type           string `json:"type" validate:"required"`
	Mount          string `json:"mount" validate:"required"`
	StartMegaBytes uint64 `json:"start_mb" validate:"required"`
	SizeMegaBytes  uint64 `json:"size_mb" validate:"required"`
	Device         string
	InstallMount   string
}
