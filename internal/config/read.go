package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}
	fmt.Printf("The file path is %v\n", path)
	file, err := os.Open(path)
	var data []byte
	size, err := file.Read(data)

	var config Config

	json.Unmarshal(data, &config)
	fmt.Printf("Size of file: %v\n", size)
	fmt.Printf("User name %v\n", config.current_user_name)
	fmt.Printf("DB URL %v\n", config.db_url)
	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homePath, configFileName)
	return fullPath, nil

}
