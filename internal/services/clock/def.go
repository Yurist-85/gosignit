package clock

import (
	"github.com/jonboulle/clockwork"
	"github.com/sarulabs/di/v2"
)

const DefinitionName = "clock"

var (
	Definition = di.Def{
		Name:  DefinitionName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return clockwork.NewRealClock(), nil
		},
	}
	DefinitionTest = di.Def{
		Name:  DefinitionName,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return clockwork.NewFakeClock(), nil
		},
	}
)
