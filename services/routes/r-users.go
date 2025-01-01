package routes

import (
	"net/http"
	"strconv"
	"wux4an/gats/services/bridges/database"

	"github.com/gin-gonic/gin"
)

func UsersTests(ctx *gin.Context) {
	nombre := ctx.Query("nombre")
	edad := ctx.Query("edad")
	if edad != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"nombre": nombre,
			"edad":   edad,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "error",
		})
	}
}

func Users(ctx *gin.Context) {
	try := ctx.Query("try")
	if try == "new" {
		name := ctx.Query("name")
		age := ctx.Query("age")
		agef, _ := strconv.ParseUint(age, 10, 32)
		database.Write(name, int(agef))
		ctx.JSON(http.StatusOK, gin.H{
			"name":    name,
			"age":     int(agef),
			"message": "Ok",
		})
	} else if try == "read" {
		name := ctx.Query("name")
		data, err := database.Read(name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"name": data,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Requests",
		})
	}
}
