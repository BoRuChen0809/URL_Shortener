package router

import (
	v1 "URL_Shortener/router/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/:id", v1.GetURL)
	r.POST("/api/v1/urls", v1.AddURL)

	return r
}
