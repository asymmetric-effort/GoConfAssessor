package utils

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYaml(fileName string, yamlObject interface{}) (err error) {
	var data []byte

	if data, err = os.ReadFile(fileName); err != nil {
		return fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	logger.Logger.Debugf("LoadYaml(): data:%q", data)

	if err = yaml.Unmarshal(data, yamlObject); err != nil {
		return fmt.Errorf("failed to unmarshal YAML %s: %w", fileName, err)
	}
	logger.Logger.Debugf("LoadYaml(): yamlObject:%v", yamlObject)
	return err
}
