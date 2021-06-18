package http

import (
	"JupiterBlog/core"
)

var E *core.Http = nil

func init() {
	E = core.Http{}.NewServer()
}

type BaseController struct{}
