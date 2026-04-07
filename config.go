package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Host  string `json:"host"`
	Title string `json:"title"`
	User  string `json:"User"`
}

func getConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(home, ".ssh_launcher", "ssh_launcher.json")
}

func loadConfigs() ([]Config, error) {
	path := getConfigPath()

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			// Return a specific error to indicate config file is missing
			return nil, fmt.Errorf("CONFIG_NOT_FOUND")
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var configs []Config
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	return configs, nil
}

func getExampleConfig() string {
	return `[
    { "host": "192.168.0.110", "title": "Mac Studio", "User": "pirogom" },
    { "host": "192.168.0.155", "title": "Mac Book Pro", "User": "pirogom" }
]`
}
