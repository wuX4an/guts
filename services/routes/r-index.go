package routes

import (
	"net/http"
	api "wux4an/gats/services/apis"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	body, err := api.Catsay()
	if err != nil {
		panic(err)
	}
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"catsay": body,
	})
}
