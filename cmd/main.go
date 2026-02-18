package main 

import "github.com/gin-gonic/gin"

func main(){

	server := gin.Default()
	server.GET("/ping", func (ctx *gin.Contex){
		ctx.JSON(200, gin.H){
			"message":"yes here we go again"
		}
	})

	server.Run(":5000")
}