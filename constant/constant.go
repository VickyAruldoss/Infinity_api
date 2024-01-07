package constants

const (
	AppConfig      = "config.json"
	ConfigFilePath = "configfiles/"
)

const (
	MR_T     string = "திரு"
	MRS_T    string = "திருமதி"
	MASTER_T string = "செல்வன்"
	SELVI    string = "செல்வி"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Others Gender = "O"
)

type MaritalStatus string

const (
	Single   MaritalStatus = "S"
	Married  MaritalStatus = "M"
	Divorced MaritalStatus = "D"
	Widow    MaritalStatus = "W"
)
