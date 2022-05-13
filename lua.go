package console

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
	"github.com/vela-security/vela-console/server"
)

var xEnv assert.Environment

func start(L *lua.LState) int {
	err := server.Instance.New(L)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	return 0
}

func WithEnv(env assert.Environment) {
	server.Instance.Inject(env)
	env.Set("console", lua.NewFunction(start))
}
