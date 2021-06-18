package core

import (
	_ "fmt"
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/xlog"
)

type engine struct {
	jupiter.Application
	//f []func() error
}

func (e engine) newEngine(f ...func() error) *engine {
	eng := &engine{}
	if err := eng.Startup(
		f...,
	); err != nil {
		xlog.Panic("startup", xlog.Any("err", err))
	}
	return eng
}

func (e *engine) Start() {
	if err := e.Run(); err != nil {
		xlog.Panic(err.Error())
	}
}
