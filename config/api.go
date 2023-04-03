package config

import (
	"github.com/lowl11/lazyconfig/confapi"
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

func SetEnvironmentName(name string) {
	_environmentName = name
}

func SetEnvironmentDefault(name string) {
	_environmentDefault = name
}

func SetEnvironmentFileName(fileName string) {
	_environmentFileName = fileName
}
