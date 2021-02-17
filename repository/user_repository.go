package repository

import "github.com/hananloser/rest-fiber-mongo-2/entity"

type UserRepository interface {
	Insert(user entity.User)

	GetAll() (user []entity.User)

	Find(Id string) (user entity.User)

	DeleteAll()
}
