package model

type Config interface {
	GetEnvironment() string
	GetPort() string
}

func New() *ConfigData {
	return &ConfigData{}
}

type ConfigData struct {
	Environment string   `json:"environment" validate:"required"`
	Port        string   `json:"port" validate:"required"`
	Postgres    Postgres `json:"postgres" validate:"required"`
}

type Postgres struct {
	Sql      string `json:"sql" validate:"required"`
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Db       string `json:"db" validate:"required"`
}

func (c *ConfigData) GetConnectionString() string {
	return c.Postgres.Sql + "://" + c.Postgres.User + ":" + c.Postgres.Password + "@" + c.Postgres.Host + "/" + c.Postgres.Db + "?sslmode=disable"
}
