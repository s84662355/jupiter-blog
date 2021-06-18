package core

import (
	_ "context"
	_ "github.com/coreos/etcd/clientv3"
	"github.com/douyu/jupiter/pkg/client/etcdv3"
	"github.com/douyu/jupiter/pkg/client/grpc/resolver"
	compound_registry "github.com/douyu/jupiter/pkg/registry/compound"
	etcdv3_registry "github.com/douyu/jupiter/pkg/registry/etcdv3"
)

type Etcd struct {
	*etcdv3.Client
	config string
}

func (e Etcd) New(config string) *Etcd {
	ee := &Etcd{
		Client: etcdv3.StdConfig(config).Build(),
		config: config,
	}
	return ee
}

func (e *Etcd) RegisterService() {
	eng.SetRegistry(
		compound_registry.New(
			etcdv3_registry.StdConfig("wh").Build(),
		),
	)
}

func (e *Etcd) RegisterClient() {
	resolver.Register("etcd", etcdv3_registry.StdConfig("wh").Build())
}
