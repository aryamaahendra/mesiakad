package middleware

import (
	"errors"
	"strings"

	"github.com/aryamaahendra/mesiakad/pkgs/api/types"
	"github.com/aryamaahendra/mesiakad/pkgs/meerrors"
	"github.com/aryamaahendra/mesiakad/pkgs/utils"
	userservice "github.com/aryamaahendra/mesiakad/services/user_service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Middleware struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Middleware {
	return &Middleware{db: db}
}

func (m *Middleware) OnlyJSON(c *fiber.Ctx) error {
	isMethodAllowed := c.Method() == "POST" || c.Method() == "PUT"
	isContentTypeAllowed := c.Get("Content-Type") != "application/json"

	if isMethodAllowed && isContentTypeAllowed {
		return c.Status(fiber.StatusUnsupportedMediaType).JSON(types.Response{
			Error:   true,
			Message: "Unsupported Media Type, only 'application/json' is accepted",
		})
	}

	return c.Next()
}

func (m *Middleware) OnlyAuthorized(c *fiber.Ctx) error {
	authorized := strings.Split(c.Get("Authorization"), " ")
	if len(authorized) < 2 {
		return meerrors.ErrUnuthorized
	}

	token := authorized[1]
	username, err := utils.ParseToken(token)
	if err != nil {
		return err
	}

	user := new(userservice.User)
	err = m.db.Where("username = ?", username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return meerrors.ErrUnuthorized
		}

		return err
	}

	c.Locals("user", user)

	return c.Next()
}
