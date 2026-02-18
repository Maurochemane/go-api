package repository



type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository{
	return ProductRepository{
		connection: connection,
	}
}

func (pr * ProductRepository) GetProducts() ([]model.Product, err){

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil{
		fmt.Println(err)
		return []model.product{}, err
	}

	var productList []model.product
	var productObj model.product

	for rows.Next(){
		err = rpws.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.price)

		if err != nil{
			fmt.Println(err)
			return []model.product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO product" + 
		"(product_name, price)" +
		" VALUES ($1, $) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return  0, err
	}

	err = query.QueryRoe(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return  0, err
	}

	query.Close()
	return id, nil
}