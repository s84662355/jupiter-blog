package core

import (
	"github.com/douyu/jupiter/pkg/server/xecho"
)

type Http struct {
	*xecho.Server
	eng *engine
}

func (h Http) NewServer(f ...func() error) *Http {

	xs := xecho.StdConfig("http").Build()

	hs := &Http{
		Server: xs,
		eng:    eng,
	}
	hs.eng.Serve(xs)
	return hs
}

func (h *Http) Run() {
	h.eng.Start()
}
