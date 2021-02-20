package repository

import "github.com/hananloser/rest-fiber-mongo-2/entity"

type ProductRepository interface {
	Insert(product entity.Product)

	GetProduct(products []entity.Product)

	Find(id int)

	DeleteAll()
}