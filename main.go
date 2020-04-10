package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/article"
	"github.com/tanimutomo/simple-api-server-go/auth"
	"github.com/tanimutomo/simple-api-server-go/handler"
)

func main() {
	articles := article.New()
	// users := user.New()

	r := gin.Default()
	r.GET("/auth", auth.GetToken())
	r.GET("/articles", handler.ArticleGet(articles))
	r.POST("/articles", handler.ArticlePost(articles))

	r.Run()
}
