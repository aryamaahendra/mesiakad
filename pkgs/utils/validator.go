package utils

import (
	"fmt"

	"github.com/aryamaahendra/mesiakad/pkgs/api/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Validator struct {
	validator *validator.Validate
}

type ErrorResponse struct {
	Error   bool   `json:"-"`
	Tag     string `json:"tag"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (mv *Validator) Validate(data interface{}) []ErrorResponse {
	validationErrs := []ErrorResponse{}

	errs := mv.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationErrs = append(validationErrs, ErrorResponse{
				Error: true,
				Tag:   err.Tag(),
				Field: err.Field(),
				Message: fmt.Sprintf(
					"validation for %s failed on the '%s' tag",
					err.Field(),
					err.Tag(),
				),
			})
		}
	}

	return validationErrs
}

/*
ParseAndValidateRequest parses JSON from the request and validates it.
If parsing or validation fails, it sends an appropriate JSON response and returns true.
Returns false if parsing and validation are successful.
*/
func ParseAndValidateRequest(c *fiber.Ctx, data any, validator *Validator) (bool, error) {
	if err := c.BodyParser(data); err != nil {
		err = c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Error:   true,
			Message: "Invalid JSON format",
		})

		if err != nil {
			return true, err
		}

		return true, nil
	}

	if errs := validator.Validate(data); len(errs) > 0 && errs[0].Error {
		err := c.Status(fiber.StatusUnprocessableEntity).JSON(types.Response{
			Error:   true,
			Message: "Validation error",
			Data:    errs,
		})

		if err != nil {
			return true, err
		}

		return true, nil
	}

	return false, nil
}
