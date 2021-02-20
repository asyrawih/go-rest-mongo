package product

import "github.com/hananloser/rest-fiber-mongo-2/model"

// In Must Implements As Interface

type ServiceProduct interface {
	// Inserting Product
	// return model.CreateProductResponse
	CreateProduct(product model.CreateProductRequest) model.CreateProductResponse

	FindProduct()

	DeleteProduct(id int)
}