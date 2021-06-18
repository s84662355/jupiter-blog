package core

import (
	"github.com/douyu/jupiter/pkg/client/redis"
)

type Redis struct {
	*redis.Redis
	config string
}

func (r Redis) New(config string) *Redis {
	cc := redis.StdRedisStubConfig(config).Build()

	rrr := &Redis{
		Redis:  cc,
		config: config,
	}
	return rrr
}
