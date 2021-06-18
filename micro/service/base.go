package service

import (
	"JupiterBlog/lib/etcd"
)

func init() {
	etcd.Client.RegisterService()
}
