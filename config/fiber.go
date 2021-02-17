package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
