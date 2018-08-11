package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/urbn/ordernumbergenerator/app"
)

var (
	GetConfigProc = GetConfig
)

func GetConfig(prefix string, Config *app.Specification) error {
	return envconfig.Process(prefix, Config)
}

func LoadConfig() (*app.Specification, error) {
	var Config app.Specification
	err := GetConfigProc("", &Config)
	return &Config, err
}
