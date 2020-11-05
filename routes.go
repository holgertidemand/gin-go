// routes.go

package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

  // Handle the index route
	router.GET("/", showIndexPage)
	
	router.GET("/article/view/:articleId", getArticle)
}



