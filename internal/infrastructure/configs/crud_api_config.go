package configs

import (
	"errors"

	"github.com/spf13/viper"
)

type CrudAPIConfig struct {
	PostgresConfig *PostgresConfig
	AppName        string `mapstructure:"app_name"`
	Port           uint16 `mapstructure:"port"`
}

func ReadCrudAPIConfig(crudAPIConfigPath string, envPath string) (*CrudAPIConfig, error) {
	var errMsg error

	pgCfg, errPgCfgReading := readPostgresConfig(envPath)
	if errPgCfgReading != nil {
		return nil, errPgCfgReading
	}

	viper.SetConfigFile(crudAPIConfigPath)
	viper.SetConfigType("yaml")

	if errCrudAPICfgReading := viper.ReadInConfig(); errCrudAPICfgReading != nil {
		return nil, errCrudAPICfgReading
	}

	var crudAPICfg CrudAPIConfig
	if errCrudAPICfgReading := viper.Unmarshal(&crudAPICfg); errCrudAPICfgReading != nil {
		return nil, errCrudAPICfgReading
	}

	if crudAPICfg.AppName == "" {
		errMsg = errors.Join(errMsg, errors.New("app_name cannot be empty"))
	}

	if crudAPICfg.Port <= 0 {
		errMsg = errors.Join(errMsg, errors.New("port must be greater than 0"))
	}

	if errMsg != nil {
		return nil, errMsg
	}

	return &CrudAPIConfig{
		PostgresConfig: pgCfg,
		AppName:        crudAPICfg.AppName,
		Port:           crudAPICfg.Port,
	}, nil
}
