package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	model "github.com/infinity-api/model"
)

type ConfigLoader interface {
	Load(fileName string, configFilePath string) (model.ConfigData, error)
}

type configurationLoader struct {
}

func NewConfigurationLoader() ConfigLoader {
	return &configurationLoader{}
}

func (c configurationLoader) Load(fileName string, configFilePath string) (model.ConfigData, error) {
	var configData model.ConfigData
	jsonfile, jsonOpenError := os.Open(filepath.Join(configFilePath, fileName))

	if jsonOpenError != nil {
		return configData, jsonOpenError
	}
	defer func(jsonfile *os.File) {
		err := jsonfile.Close()
		if err != nil {
		}
	}(jsonfile)

	byteValue, jsonReadError := ioutil.ReadAll(jsonfile)
	if jsonReadError != nil {
		return configData, jsonReadError
	}
	err := json.Unmarshal(byteValue, &configData)
	if err != nil {
		return model.ConfigData{}, err
	}
	return configData, nil
}
