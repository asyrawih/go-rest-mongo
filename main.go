package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hananloser/rest-fiber-mongo-2/config"
	"github.com/hananloser/rest-fiber-mongo-2/controller"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
	"github.com/hananloser/rest-fiber-mongo-2/repository"
	"github.com/hananloser/rest-fiber-mongo-2/service/user"
)

func main()  {
	// Setup Database
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// Setup repository
	userRepository := repository.NewUserRepository(database)
	//For ignore variable
	_ = repository.NewProductRepository(database)
	// Setup Service
	userService := service_user.NewUserService(&userRepository)
	//Setup Controller
	userController := controller.NewUserController(userService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())

	// Register The Route
	userController.Route(app)

	// Running THE APP
	err := app.Listen(":8000")
	exception.PanicIfNeeded(err)
}