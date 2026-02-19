

# instalar bibliotecas
go mod tidy

# Lib Http
go get github.com/gin-gonic/gin 

# Lib Postgres
_ "github.com/lib/pq"

# Testando a lib

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

# Subindo o container postgres
--docker compose up -d go_db

# Mockando um dado

` func (p *productController) GetProducts(ctx *gin.Contex){

	products := []model.product{
		{
			ID: 1,
			Name: "Nanami",
			Price: 200,
		}
	}

	ctx.JSON(http.StatusOk, products)
}`

# Build da imagem

docker build -t go-api .