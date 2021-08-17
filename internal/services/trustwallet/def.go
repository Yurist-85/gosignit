package trustwallet

import (
	"github.com/sarulabs/di/v2"
)

const DefinitionName = "trustwallet"

var Definition = di.Def{
	Name:  DefinitionName,
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		return NewSigner(), nil
	},
}
