package article

import (
	"JupiterBlog/core"
	"JupiterBlog/micro/protobuf/service"
	_ "JupiterBlog/micro/service"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
	"github.com/douyu/jupiter/pkg/xlog"
)

func init() {
	server := xgrpc.StdConfig("grpc").Build()

	service.RegisterArticleServer(server.Server, &Article{
		server: server,
	})
	eng := core.GetEngine()
	eng.Serve(server)
}

func Start() {
	eng := core.GetEngine()
	if err := eng.Run(); err != nil {
		xlog.Error(err.Error())
	}
}
