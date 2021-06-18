package article

import (
	_ "JupiterBlog/micro/client"
	"JupiterBlog/micro/protobuf/service"
	"github.com/douyu/jupiter/pkg/client/grpc"
	"github.com/douyu/jupiter/pkg/client/grpc/balancer"
)

var articleClient service.ArticleClient

func init() {
	config := grpc.StdConfig("article")
	config.BalancerName = balancer.NameSmoothWeightRoundRobin
	articleClient = service.NewArticleClient(config.Build())
}
