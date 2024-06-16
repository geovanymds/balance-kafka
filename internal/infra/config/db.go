package config

import "github.com/geovanymds/balance/pkg/utils"

type DbConnectionConfig struct {
	Dialect  string
	Port     string
	User     string
	Password string
	Host     string
	DbName   string
}

func NewDbConnectionConfig() *DbConnectionConfig {
	return &DbConnectionConfig{
		utils.GetEnvWithDefaultValue("DB_DIALECT", "mysql"),
		utils.GetEnvWithDefaultValue("DB_PORT", "3306"),
		utils.GetEnvWithDefaultValue("DB_USER", "root"),
		utils.GetEnvWithDefaultValue("DB_PASSWORD", "root"),
		utils.GetEnvWithDefaultValue("DB_HOST", "mysql"),
		utils.GetEnvWithDefaultValue("DB_NAME", "wallet"),
	}
}
