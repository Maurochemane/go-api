package main 

import "github.com/gin-gonic/gin"

func main(){

	server := gin.Default()

	//conect database
	dbConnection, err := db.ConnectDB()
	if err != nil{
		panic(err)
	}


	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)


	// camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)


	//camada de controllers
	productController := controller.NewProductController(ProductUsecase)






	server.GET("/ping", func (ctx *gin.Contex){
		ctx.JSON(200, gin.H){
			"message":"yes here we go again"
		}
	})


	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.Get("/product/:productId", productController.GetProductById)
	//PUT
	//DELETE

	//JWt auth

	server.Run(":5000")
}