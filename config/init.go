package config

import "github.com/spf13/viper"

const configFilePath string = ".env"

func NewAppConfig() (*AppConfig, error) {
	v, err := initViper(configFilePath)
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Postgres: Postgres{
			Host: v.GetString("DB_HOST"),
			Port: v.GetString("DB_PORT"),
			User: v.GetString("DB_USER"),
			Pass: v.GetString("DB_PASSWORD"),
			Name:v.GetString("DB_NAME"),
		},
		Fiber: Fiber{
			Port: v.GetString("FIBER_PORT"),
		},
	}, nil
}

func initViper(configFilePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFilePath)

	err := v.ReadInConfig()
	if err != nil {
		return nil, err 
	}
	
	return v, nil
}
