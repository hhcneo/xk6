package sproto_sprotodefine

import (
	"github.com/traefik/yaegi/interp"
	"go.k6.io/k6/js/modules"
)

var Symbols = interp.Exports{}

func Exports(vu modules.VU) interp.Exports {
	return Symbols
}
