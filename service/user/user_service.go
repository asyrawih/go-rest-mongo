package service_user

import "github.com/hananloser/rest-fiber-mongo-2/model"

type UserService interface {
	 // Create User
	 Create(request  model.CreateUserRequest) (response model.CreateUserResponse)
	 // List User
	 List() (response []model.GetUserResponse)
}
