package client

import (
	"JupiterBlog/lib/etcd"
)

func init() {
	etcd.Client.RegisterClient()
}
