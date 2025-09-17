package order

type Product struct {
	ProductId   int64  `json:"productId"`
	ProductName string `json:"productName"`
	Price       int64  `json:"price"`
}

type ProductService struct {
}

func (p *ProductService) GetProductById(productId int64) (*Product, error) {
	return &Product{ProductId: productId, ProductName: "DELL显示器", Price: 2185}, nil
}
