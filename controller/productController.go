package controller

import (
	"go-api/model"
	"go-api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService service.ProductService
}

func NewProductController(service service.ProductService) productController {
	return productController{
		productService: service,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product
	// Pega o JSON e transforma num product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productService.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Reponse{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Reponse{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productService.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Reponse{
			Message: "Produto não foi encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Reponse{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Reponse{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		response := model.Reponse{
			Message: "Falha ao ler o corpo da requisição, verifique o tipo dos atributos: " + err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Só aqui fazemos as validações de campo porque o bind deu certo
	if product.Name == "" {
		response := model.Reponse{
			Message: "É obrigatório informar o nome do produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if product.Price <= 0 {
		response := model.Reponse{
			Message: "O preço precisa ser positivo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product.ID = productId
	log.Println("Recebido no controller:", product)

	updatedProduct, err := p.productService.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Reponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProduct)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("productId")
	if id == "" {
		response := model.Reponse{
			Message: "ID do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Reponse{
			Message: "ID do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = p.productService.DeleteProduct(productId)
	if err != nil {
		response := model.Reponse{
			Message: err.Error(),
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := model.Reponse{
		Message: "Sucesso ao excluir registro",
	}
	ctx.JSON(http.StatusOK, response)
}
