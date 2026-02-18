package controller

import(
	"go-api/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type productController struct {
	//usecase
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase ) productController{
	return productController{
		productUsecase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Contex){

	products, err := p.productUsecase.GetProducts()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOk, products)
}

func (p * productController) CreateProduct(ctx.Contex){

	var product model.product
	err := ctx.BindJSON(&product)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productCase.CreateProduct(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.json(http.StatusCreated, insertedProduct)
}