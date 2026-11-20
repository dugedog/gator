package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	stats, err := os.Stat(path)
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	data := make([]byte, stats.Size())
	_, err = file.Read(data)
	if err = file.Close(); err != nil {
		return Config{}, nil
	}
	var configuration Config

	if err = json.Unmarshal(data, &configuration); err != nil {
		return Config{}, err

	}
	return configuration, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(homePath, configFileName)
	return fullPath, nil

}
