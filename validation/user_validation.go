package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
	"github.com/hananloser/rest-fiber-mongo-2/model"
)

func Validate(request model.CreateUserRequest) {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Address, validation.Required),
		validation.Field(&request.PhoneNumber, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
