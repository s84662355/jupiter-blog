package article

import (
	"JupiterBlog/micro/protobuf/request"
	"JupiterBlog/micro/protobuf/response"
	"JupiterBlog/models"
	"context"
	"fmt"
	"github.com/douyu/jupiter/pkg/server/xgrpc"
)

type Article struct {
	server *xgrpc.Server
}

func (g Article) Info(context context.Context, req *request.ArticleInfo) (*response.ArticleInfo, error) {

	Article := &response.ArticleInfo{}
	models.Article{}.Model().Table(models.Article{}.TableName()).Where("id = 1").First(Article)
	fmt.Println(Article)
	return Article, nil
}
