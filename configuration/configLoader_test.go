package configuration

import (
	"testing"

	model "github.com/infinity-api/model"
	"github.com/stretchr/testify/suite"
)

type ConfigurationLoaderTestSuite struct {
	suite.Suite
	ConfigLoader
}

func TestConfigurationLoaderTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigurationLoaderTestSuite))
}

func (suite *ConfigurationLoaderTestSuite) SetupTest() {
	suite.ConfigLoader = NewConfigurationLoader()
}

func (s *ConfigurationLoaderTestSuite) TestShouldCallLoadTheConfigurationFromConfigJSON() {
	expectedConfig := model.ConfigData{
		Environment: "DEV",
		Port:        "3000",
		Postgres: model.Postgres{
			Sql:      "PostgreSQL",
			Host:     "localhost",
			User:     "postgres",
			Password: "admin",
			Db:       "todos",
		},
	}
	expectedConnectionString := "PostgreSQL://postgres:admin@localhost/todos?sslmode=disable"
	data, _ := s.ConfigLoader.Load("testconfig.json", "../configfiles/")
	connectionString := data.GetConnectionString()
	s.Suite.Equal(expectedConfig, data)
	s.Suite.Equal(expectedConnectionString, connectionString)
}

func (s *ConfigurationLoaderTestSuite) TestShouldThrowErrorWhenTheFileDoesnotExists() {
	_, err := s.ConfigLoader.Load("test.json", "../configfiles/")
	s.Suite.NotNil(err)
}
