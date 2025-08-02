package service

import (
	"go-api/model"
	"go-api/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return ProductService{
		repository: repo,
	}
}

func (ps *ProductService) GetProducts() ([]model.Product, error) {
	return ps.repository.GetProducts()
}

func (ps *ProductService) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := ps.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (ps *ProductService) GetProductById(idProdut int) (*model.Product, error) {
	product, err := ps.repository.GetProductById(idProdut)
	if err != nil {
		return nil, err
	}

	return product, nil
}
