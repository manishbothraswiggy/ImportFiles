package data

type Product struct{
	ID int
	Name string
	Description string
	Price float64
	SKU string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

func GetProducts() []*Product{
	return productList
}
var productList = []*Product{
	&Product{
		ID: 1,
		Name: Product1,
		Description: Product1,
		Price: 111.11,
		SKU: SKU1,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Product2",
		Description: Product2,
		Price: 222.22,
		SKU: SKU2,
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
		DeletedOn: time.Now().UTC().String(),
	}
}