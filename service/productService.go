package service

import (
	"go-api/model"
	"go-api/repository"
	"log"
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

func (ps *ProductService) UpdateProduct(product model.Product) (*model.Product, error) {
	log.Println("Recebido na service:", product)

	updatedProduct, err := ps.repository.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (ps *ProductService) DeleteProduct(idProdut int) error {
	err := ps.repository.DeleteProduct(idProdut)
	if err != nil {
		return err
	}

	return nil
}
