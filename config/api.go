package config

import (
	"github.com/lowl11/lazyconfig/confapi"
	"github.com/lowl11/lazyconfig/config_data"
	"os"
	"strings"
)

func Init() {
	if err := confapi.NewMap().
		EnvironmentDefault(_environmentDefault).
		EnvironmentName(_environmentName).
		EnvFileName(_environmentFileName).
		Read(); err != nil {
		panic("Read map config error: " + err.Error())
	}
}

func Get(key string) string {
	return confapi.Get(key)
}

func Env() string {
	if _environment == "" {
		_environment = os.Getenv("env")
	}

	return strings.ToLower(_environment)
}

func IsProduction() bool {
	return strings.ToLower(Env()) == "production"
}

func IsLocal() bool {
	return strings.ToLower(Env()) == "local"
}

func SetEnvironmentName(name string) {
	if name == "" {
		_environmentName = config_data.EnvironmentName
		return
	}
	_environmentName = name
}

func SetEnvironmentDefault(name string) {
	if name == "" {
		_environmentDefault = config_data.EnvironmentDefault
		return
	}
	_environmentDefault = name
}

func SetEnvironmentFileName(fileName string) {
	if fileName == "" {
		_environmentFileName = config_data.EnvFileNameDefault
		return
	}
	_environmentFileName = fileName
}
