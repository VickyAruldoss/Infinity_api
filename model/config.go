package model

type Config interface {
	GetEnvironment() string
	GetPort() string
}

func New() *ConfigData {
	return &ConfigData{}
}

type ConfigData struct {
	Environment string `json:"environment" validate:"required"`
	Port        string `json:"port" validate:"required"`
}
