package trustwallet

import (
	"github.com/sarulabs/di/v2"

	"github.com/yurist-85/gosignit/internal/config"
)

const DefinitionName = "trustwallet"

var Definition = di.Def{
	Name:  DefinitionName,
	Scope: di.App,
	Build: func(ctn di.Container) (interface{}, error) {
		cfg := ctn.Get(config.DefinitionName).(config.ConfigInterface)

		return NewTrustWallet(cfg), nil
	},
}
