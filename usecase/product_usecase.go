package usecase


type ProductUsecase struct {
	//repository
	repository repository.ProductRepository
}


func NewProductUsecase() ProductUsecase{
	return ProductUsecase{
		repository: repo
	}
}

func ( pu *ProductUsecase) GetProducts() ([]model.Product, error){
	return pu.repository.GetProducts()
}

func(pu *ProductUsecase) CreateProduct(product model.product) (model.product, err){

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId
	return product, nil
}