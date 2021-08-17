package config

import "github.com/sarulabs/di/v2"

const DefinitionName = "config"

var Definition = di.Def{
	Name: DefinitionName,
	Build: func(ctn di.Container) (i interface{}, err error) {
		c := NewConfig()

		return c, nil
	},
}
