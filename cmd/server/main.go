package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


const Port string = ":42069"

func main() {
	router := gin.Default()
	
	router.Static("/static", "./web/static/")
	router.LoadHTMLGlob("web/templates/*")


	router.GET("/make", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "post.html", gin.H{})
	})


	router.Run(Port)
}
