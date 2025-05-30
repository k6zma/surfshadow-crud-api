package configs

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Host     string `env:"PG_HOST" envDefault:"localhost"`
	Port     uint16 `env:"PG_PORT" envDefault:"5432"`
	User     string `env:"PG_USER" envDefault:"postgres"`
	Password string `env:"PG_PASSWORD" envDefault:"password"`
	Name     string `env:"PG_NAME" envDefault:"postgres"`
}

func readPostgresConfig(envPath string) (*PostgresConfig, error) {
	var errMsg error

	errLoadEnv := godotenv.Overload(envPath)
	if errLoadEnv != nil {
		return nil, errLoadEnv
	}

	pgHost := os.Getenv("PG_HOST")
	if pgHost == "" {
		errMsg = errors.Join(errMsg, errors.New("PG_HOST cannot be empty"))
	}

	pgPort := os.Getenv("PG_PORT")
	if pgPort == "" {
		errMsg = errors.Join(errMsg, errors.New("PG_PORT cannot be empty"))
	}

	pgPortUint, err := strconv.ParseUint(pgPort, 10, 16)
	if err != nil {
		errMsg = errors.Join(errMsg, errors.New("PG_PORT must be a valid number"))
	}

	pgUsername := os.Getenv("PG_USER")
	if pgUsername == "" {
		errMsg = errors.Join(errMsg, errors.New("PG_USER cannot be empty"))
	}

	pgPassword := os.Getenv("PG_PASSWORD")
	if pgPassword == "" {
		errMsg = errors.Join(errMsg, errors.New("PG_PASSWORD cannot be empty"))
	}

	pgName := os.Getenv("PG_NAME")
	if pgName == "" {
		errMsg = errors.Join(errMsg, errors.New("PG_NAME cannot be empty"))
	}

	if errMsg != nil {
		return nil, errMsg
	}

	return &PostgresConfig{
		Host:     pgHost,
		Port:     uint16(pgPortUint),
		User:     pgUsername,
		Password: pgPassword,
		Name:     pgName,
	}, nil
}
