package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Endpoint string `env:"CLAI_ENDPOINT" cli:"endpoint" usage:"The API endpoint to use"`
	Token    string `env:"CLAI_TOKEN" cli:"token" usage:"The API token to use"`
	Model    string `env:"CLAI_MODEL" cli:"model" usage:"The model to use"`
	Shell    string `env:"CLAI_SHELL" cli:"shell" usage:"The target shell to generate commands for, empty for current shell"`
	OS       string `env:"CLAI_OS" cli:"os" usage:"The target OS to generate commands for, empty for current OS"`
	Phrase   string `arg:"1"`
}

func FindConfigFile() (string, error) {
	filename := ".clai.yaml"
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(filepath.Join(home, filename)); err == nil {
		return filepath.Join(home, filename), nil
	}

	return "", fmt.Errorf("failed to find config file %s", filename)
}
