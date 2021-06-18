package article

import (
	"JupiterBlog/lib/helper"
	"JupiterBlog/micro/protobuf/request"
	"JupiterBlog/micro/protobuf/response"
	_ "JupiterBlog/micro/protobuf/service"
	"context"
)

type Article struct {
	Name string
}

func (m Article) Info(id interface{}) *response.ArticleInfo {
	req := &request.ArticleInfo{}
	req.Id = helper.Str2Uint64(helper.Int2Str(id))
	resp, err := articleClient.Info(context.Background(), req)

	if err != nil {
		return nil
	}
	return resp
}
