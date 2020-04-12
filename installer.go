package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

type InstallConfig struct {
	Version      string             `json:"version"`
	Keyboard     KeyboardConfig     `json:"keyboard"`
	Localization LocalizationConfig `json:"localization"`
	Computer     ComputerConfig     `json:"computer"`
	Storage      StorageConfig      `json:"storage"`
}

type KeyboardConfig struct {
	Layout string `json:"layout"`
}

type LocalizationConfig struct {
	Language string   `json:"language"`
	Locales  []string `json:"locales"`
	Timezone string   `json:"timezone"`
}

type ComputerConfig struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"Password"`
}

type StorageConfig struct {
	Devices []DeviceConfig `json:"devices"`
}

type DeviceConfig struct {
	Device     string            `json:"device"`
	Partitions []PartitionConfig `json:"partitions"`
}

type PartitionConfig struct {
	Type  string `json:"type"`
	Mount string `json:"mount"`
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

func ReadConfig(bytes []byte, config *InstallConfig) error {
	return json.Unmarshal(bytes, config)
}
