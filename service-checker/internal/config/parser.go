package config

import (
	"encoding/json"
	"fmt"
	"service-checker/internal/model"
)

func ParseConfig(filepath string) ([]model.Target, error) {
	configByte, err := readConfigFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to Read Config File: %w", err)
	}

	var targets []model.Target
	err = json.Unmarshal(configByte, &targets)
	if err != nil {
		return nil, fmt.Errorf("Failed to Unmarshal JSON: %w", err)
	}

	return targets, err
}
