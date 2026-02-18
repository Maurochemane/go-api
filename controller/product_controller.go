package controller

import(
	"go-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productController struct {
	//usecase
}

func NewProductController(){
	return productController
}

func (p *productController) GetProducts(ctx *gin.Contex){

	products := []model.product{
		{
			ID: 1,
			Name: "Nanami",
			Price: 200,
		}
	}

	ctx.JSON(http.StatusOk, products)
}