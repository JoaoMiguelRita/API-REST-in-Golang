package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
	"log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, productName, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()
	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := pr.connection.Prepare("INSERT INTO product(productName, price) VALUES($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(idProduct int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(idProduct).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &produto, nil
}

func (pr *ProductRepository) UpdateProduct(product model.Product) (*model.Product, error) {
	log.Println("Recebido no repository:", product)

	query, err := pr.connection.Prepare("UPDATE product SET productName = $2, price = $3 WHERE id = $1")
	if err != nil {
		return nil, err
	}

	result, err := query.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("Nenhuma linha foi atualizada, verifique se o produto com id %d não exista", product.ID)
	}
	query.Close()
	return &product, nil
}

func (pr *ProductRepository) DeleteProduct(idProduct int) error {
	log.Println("Recebido no repository idProduct:", idProduct)

	query, err := pr.connection.Prepare("DELETE FROM product WHERE id = $1")
	if err != nil {
		return err
	}

	result, err := query.Exec(idProduct)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %d não encontrado", idProduct)
	}

	query.Close()
	return nil
}
