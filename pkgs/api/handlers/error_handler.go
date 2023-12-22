package handlers

import (
	"errors"

	"github.com/aryamaahendra/mesiakad/pkgs/api/types"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		err = ctx.Status(code).JSON(types.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	if errors.Is(err, meerrors.ErrRecordNotFound) {
		err = ctx.Status(fiber.StatusNotFound).JSON(types.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	if errors.Is(err, meerrors.ErrUnuthorized) {
		err = ctx.Status(fiber.StatusUnauthorized).JSON(types.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	if _, ok := err.(*meerrors.ErrNotUniqueInDB); ok {
		err = ctx.Status(fiber.StatusUnprocessableEntity).JSON(types.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	if err != nil {
		// In case the SendFile fails
		return ctx.Status(fiber.StatusInternalServerError).JSON(types.Response{
			Error:   true,
			Message: err.Error(),
		})
	}

	return nil
}
