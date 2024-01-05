package main

import (
	"os"

	cfg "github.com/crslex/miniProject/config"
	"gopkg.in/yaml.v3"
)

func InitConfig(configPath string) (*cfg.Config, error) {
	config := &cfg.Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// yaml decode
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil

}
