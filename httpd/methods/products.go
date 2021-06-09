package handler

type Product struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Stock       string `json:"stock"`
}

type Repository struct {
	Products []Product
}

func New() *Repository {
	return &Repository{
		Products: []Product{},
	}
}

func (repo *Repository) Add(product Product) {
	repo.Products = append(repo.Products, product)
}

func (repo *Repository) GetAll() []Product {
	return repo.Products
}
