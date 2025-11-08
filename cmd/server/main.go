package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


const Port string = ":42069"

func main() {
	router := gin.Default()


	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "dan smells")
	})
	router.Run(Port)
}
