package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
	"github.com/hananloser/rest-fiber-mongo-2/model"
	services "github.com/hananloser/rest-fiber-mongo-2/service/user"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{UserService: userService}
}


func (controller *UserController) Route(app *fiber.App){
	app.Post("/api/users" , controller.Create)
	app.Get("/api/users" ,  controller.GetAll)
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := ctx.BodyParser(&request)

	request.Id = uuid.New().String()
	exception.PanicIfNeeded(err)

	response := controller.UserService.Create(request)

	return ctx.JSON(model.WebResponse{
		Code:    200,
		Message: "Data Has Been Created",
		Data:    response,
	})
}

func (controller UserController) GetAll(ctx *fiber.Ctx) error {
	var response = controller.UserService.List()
	return ctx.JSON(model.WebResponse{
		Code:    200,
		Message: "GOT",
		Data:    &response,
	})
}


