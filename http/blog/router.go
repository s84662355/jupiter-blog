package blog

import (
	"JupiterBlog/http"
	"JupiterBlog/micro/client/article"
	"JupiterBlog/models"
	"github.com/labstack/echo/v4"
)

func init() {
	http.E.GET("/hello", func(ctx echo.Context) error {
		res := []*models.Cate{}
		models.Cate{}.Model().Scan(&res)

		return ctx.JSON(200, article.Article{}.Info(1))
	})
}

func Start() {
	http.E.Run()
}
