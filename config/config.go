package config

import "github.com/realHoangHai/gmeet-biz/utils"

type Config struct {
	Database *DatabaseConfig
	Jwt      *JwtConfig
}

func New() *Config {
	return &Config{
		Database: NewDatabaseConfig(),
		Jwt:      NewJwtConfig(),
	}
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     utils.GetIni("database", "HOST", "localhost"),
		Name:     utils.GetIni("database", "DATABASE_NAME", "gmeet"),
		User:     utils.GetIni("database", "DATABASE_USER", "postgres"),
		Password: utils.GetIni("database", "DATABSE_PASSWORD", "postgres"),
		Port:     utils.GetIni("database", "DATABSE_PORT", "5432"),
	}
}

type JwtConfig struct {
	Secret []byte
}

func NewJwtConfig() *JwtConfig {
	return &JwtConfig{
		Secret: []byte(utils.GetIni("jwt", "SECRET", "secret")),
	}
}
