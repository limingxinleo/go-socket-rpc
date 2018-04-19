package tests

import (
	"github.com/limingxinleo/gorpc/tests/providers"
)

type Providers struct {
	Config *providers.Config
	DB     *providers.Mysql
}

func GetProviders() *Providers {
	config := &providers.Config{}
	config.Register()

	mysql := &providers.Mysql{}
	mysql.Register()

	return &Providers{
		Config: config,
		DB:     mysql,
	}
}

var Provider = GetProviders()
