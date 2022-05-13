package server

import (
	"github.com/vela-security/vela-public/assert"
	"github.com/vela-security/vela-public/lua"
)

var xEnv assert.Environment

func (s *Serv) Inject(env assert.Environment) {
	xEnv = env

	Instance = newServ()
	xEnv.Register(Instance)

	env.Set("print", lua.NewFunction(s.output))
}
