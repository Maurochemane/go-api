package main 

import "github.com/gin-gonic/gin"

func main(){

	server := gin.Default()

	//camada de controllers
	productController := controller.NewProductController()

	server.GET("/ping", func (ctx *gin.Contex){
		ctx.JSON(200, gin.H){
			"message":"yes here we go again"
		}
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":5000")
}