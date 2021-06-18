package redis

import (
	"JupiterBlog/core"
)

var Client *redis = nil

type Redis struct {
	*core.Redis
}

func init() {
	Client = &Redis{
		Etcd: core.Redis{}.New("default"),
	}
}
