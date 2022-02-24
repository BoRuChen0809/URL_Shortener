package router

import (
	v1 "URL_Shortener/router/api/v1"
	v2 "URL_Shortener/router/api/v2"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/mysql/:id", v1.GetURL)
	r.POST("/api/mysql/urls", v1.AddURL)

	r.GET("/:id", v2.GetURL)
	r.POST("/api/v1/urls", v2.AddURL)
	return r
}
