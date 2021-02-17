package exception

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hananloser/rest-fiber-mongo-2/model"
)

func ErrorHandler (ctx *fiber.Ctx, err error) error {
	_ , oke := err.(ValidationError)
	if oke {
		return ctx.JSON(model.WebResponse{
			Code:    400,
			Message: "BAD_REQUEST",
			Data:    err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:    500,
		Message: "INTERNAL SERVER ERROR",
		Data:    err.Error(),
	})
}