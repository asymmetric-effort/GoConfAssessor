package utils

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYaml(fileName string, yamlObject interface{}) (err error) {
	var data []byte

	if data, err = os.ReadFile(fileName); err != nil {
		return fmt.Errorf("failed to read file %s: %w", fileName, err)
	}

	if err = yaml.Unmarshal(data, &yamlObject); err != nil {
		return fmt.Errorf("failed to unmarshal YAML %s: %w", fileName, err)
	}

	return err
}
