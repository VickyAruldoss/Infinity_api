package constants

const (
	AppConfig      = "config.json"
	ConfigFilePath = "configfiles/"
)

type Gender int

const (
	Male Gender = iota + 1
	Female
	Others
)
