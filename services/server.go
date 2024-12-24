package services

import (
	"net/http"
	"wux4an/gats/services/routes"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	r := gin.New()

	r.LoadHTMLGlob("web/templates/*")
	r.Static("/static", "web/static")
	r.Static("/assets", "web/assets")
	r.StaticFile("/favicon.ico", "web/assets/favicon.png")

	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "404.html", nil)
	})

	r.GET("/", routes.Index)
	return r
}
