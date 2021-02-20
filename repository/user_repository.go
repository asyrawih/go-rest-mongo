package repository

import "github.com/hananloser/rest-fiber-mongo-2/entity"

type UserRepository interface {
	//Insert User
	Insert(user entity.User)
	// GetAll user
	GetAll(userId string) (user []entity.User)
	//Find User
	Find(Id string) (user entity.User)
	//DeleteAll User
	DeleteAll()
}
