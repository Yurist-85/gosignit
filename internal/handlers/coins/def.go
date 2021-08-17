package coins

import "github.com/sarulabs/di/v2"

const DefinitionName = "handler-coins"

var Definition = di.Def{
	Name: DefinitionName,
	Build: func(ctn di.Container) (i interface{}, err error) {
		handler := NewHandler(ctn)

		return handler, nil
	},
}
