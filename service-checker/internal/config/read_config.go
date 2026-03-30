package config

import (
	"fmt"
	"os"
)

func readConfigFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Failed to Read Config File:", err)
		return nil, err
	}
	return data, nil
}
