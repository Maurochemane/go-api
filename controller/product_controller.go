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

func (p *productController) GetProductById(ctx *gin.Contex){

	id := ctx.Param("productId")
	if (id == ""){
		response := model.Response{
			Message:"Id do producto nao pode ser nulo",
		}
		ctx.json(http.StatusBadRequest, )
		return
	}


	productId, err := strconv.Atoi(id)
	if (err != nil){
		response := model.Response{
			Message:"Id do produto precisa ser um numero",
		}
		ctx.json(http.StatusBadRequest, )
		return
	}


	product, err := p.productUsecase.GetProductById(productId)
	if (err != nil){
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message:"O producto nao foi encontrado na base de dados",
		}
		ctx.json(http.StatusNotFound,response )
		return
	}

	ctx.JSON(http.StatusOk, product)
}