package repository

import (
	"github.com/hananloser/rest-fiber-mongo-2/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryImpl struct {
	Collection *mongo.Collection
}

func (p ProductRepositoryImpl) Insert(product entity.Product) {
	panic("implement me")
}

func (p ProductRepositoryImpl) GetProduct(products []entity.Product) {
	panic("implement me")
}

func (p ProductRepositoryImpl) Find(id int) {
	panic("implement me")
}

func (p ProductRepositoryImpl) DeleteAll() {
	panic("implement me")
}

func NewProductRepository(database *mongo.Database) ProductRepository {
	return &ProductRepositoryImpl{Collection: database.Collection("product")}
}
