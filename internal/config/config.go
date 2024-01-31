package config

import "fmt"

type Config struct {
	Http struct {
		Address string `env:"HTTP_ADDRESS, default=:8080"`
	}

	Postgres struct {
		Host     string `env:"PG_HOST"`
		Port     string `env:"PG_PORT, default=5432"`
		User     string `env:"PG_USER, default=postgres"`
		Password string `env:"PG_PASS"`
		DBName   string `env:"PG_NAME, default=postgres"`
	}

	EnableDocs int `env:"ENABLE_DOCS"`
}

func New() Config {
	return Config{}
}

// postgres://username:password@host:port/database_name
const PGConnFstring = "postgres://%s:%s@%s:%s/%s"

func (c *Config) PGConnString() string {
	return fmt.Sprintf(
		PGConnFstring,
		c.Postgres.User,
		c.Postgres.Password,
		c.Postgres.Host,
		c.Postgres.Port,
		c.Postgres.DBName,
	)
}
