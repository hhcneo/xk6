package neo

import (
	"github.com/hhcneo/xk6/internal/builtin/gjson"
	"github.com/hhcneo/xk6/internal/builtin/gofakeit"
	"github.com/hhcneo/xk6/internal/builtin/goquery"
	"github.com/hhcneo/xk6/internal/builtin/jsonpath"
	"github.com/hhcneo/xk6/internal/builtin/jsonschema"
	"github.com/hhcneo/xk6/internal/builtin/logrus"
	"github.com/hhcneo/xk6/internal/builtin/resty"
	"github.com/hhcneo/xk6/internal/builtin/sproto_sprotodefine"
	"github.com/hhcneo/xk6/internal/builtin/stdlib"
	"github.com/hhcneo/xk6/internal/builtin/testify"
	"go.k6.io/k6/js/modules"
)

func registerBuiltins() {
	RegisterExports(stdlib.Exports)
	RegisterExports(logrus.Exports)
	RegisterExports(resty.Exports)
	RegisterExports(testify.Exports)
	RegisterExports(goquery.Exports)
	RegisterExports(gjson.Exports)
	RegisterExports(jsonpath.Exports)
	RegisterExports(jsonschema.Exports)
	RegisterExports(gofakeit.Exports)
	RegisterExports(sproto_sprotodefine.Exports)
}

func registerExtension() {
	modules.Register("k6/x/neo", New())
}

func Bootstrap() {
	redirectStdin()
	registerBuiltins()
	registerExtension()
}
