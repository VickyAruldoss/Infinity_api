package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	model "github.com/infinity-api/model"
)

type ConfigurationLoader interface {
	Load(fileName string, configFilePath string) (model.ConfigData, error)
}

type configurationLoader struct {
}

func NewConfigurationLoader() *configurationLoader {
	return &configurationLoader{}
}

func (c configurationLoader) Load(fileName string, configFilePath string) (model.ConfigData, error) {
	var configData model.ConfigData
	jsonfile, jsonOpenError := os.Open(filepath.Join(configFilePath, fileName))

	if jsonOpenError != nil {
		return configData, jsonOpenError
	}
	defer jsonfile.Close()

	byteValue, jsonReadError := ioutil.ReadAll(jsonfile)
	if jsonReadError != nil {
		return configData, jsonReadError
	}
	json.Unmarshal(byteValue, &configData)
	return configData, nil
}
