package config

import (
	"encoding/json"
	"fmt"
	"service-checker/internal/model"
)

func ParseConfig(filepath string) ([]model.Target, error) {
	configByte, err := readConfigFile(filepath)
	if err != nil {
		return nil, err
	}

	var targets []model.Target
	err = json.Unmarshal(configByte, &targets)
	if err != nil {
		fmt.Println("Failed to Deserialize JSON:", err)
		return nil, err
	}

	return targets, err
}
