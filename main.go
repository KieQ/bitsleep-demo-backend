package main

import (
	"demo/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	urlBind(r)

	r.Run(":8080") //default to 8080 port
}


func urlBind(r *gin.Engine){
	image := r.Group("/image")
	image.POST("/generate", handler.ImageGenerate)

}