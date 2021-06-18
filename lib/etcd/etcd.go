package etcd

import (
	"JupiterBlog/core"
)

var Client *Etcd = nil

type Etcd struct {
	*core.Etcd
}

func init() {
	Client = &Etcd{
		Etcd: core.Etcd{}.New("default"),
	}
}
