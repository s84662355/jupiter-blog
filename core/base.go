package core

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var eng *engine

func init() {
	eng = engine{}.newEngine(
		func() error { return nil },
	)
}

func GetEngine() *engine {
	return eng
}
