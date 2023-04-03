package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yosikez/article-go/handler"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handler.ShowIndexPage)
	router.GET("/article/view/:id", handler.GetArticle)

	router.Run()
}
