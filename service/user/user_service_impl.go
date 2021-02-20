package service_user

import (
	"github.com/hananloser/rest-fiber-mongo-2/entity"
	"github.com/hananloser/rest-fiber-mongo-2/model"
	"github.com/hananloser/rest-fiber-mongo-2/repository"
	"github.com/hananloser/rest-fiber-mongo-2/validation"
)

type ServiceImpl struct {
	UserRepository repository.UserRepository
}



func NewUserService(userRepository *repository.UserRepository) UserService {
	return &ServiceImpl{
		UserRepository: *userRepository,
	}
}

func (userService ServiceImpl) Create(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.Validate(request)

	user := entity.User{
		Id:          request.Id,
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	userService.UserRepository.Insert(user)

	response = model.CreateUserResponse{
		Id:          request.Id,
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
	}

	return response
}

func (userService *ServiceImpl) List(userId  string) (response []model.GetUserResponse) {
	users := userService.UserRepository.GetAll(userId)
	for _, user := range users {
		response = append(response, model.GetUserResponse{
			Id:          user.Id,
			Name:        user.Name,
			Email:       user.Email,
			Password:    user.Password,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
		})
	}
	return response
}

func (userService ServiceImpl) Find(id string) (response model.FindUserResponse) {
	var user  = userService.UserRepository.Find(id)
	response = model.FindUserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
	}

	return response

}
