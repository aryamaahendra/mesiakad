package handlers

import (
	"errors"

	"github.com/aryamaahendra/mesiakad/domains/account"
	"github.com/aryamaahendra/mesiakad/pkgs/api/types"
	"github.com/aryamaahendra/mesiakad/pkgs/utils"
	userservice "github.com/aryamaahendra/mesiakad/services/user_service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	user      *userservice.UserService
	validator *utils.Validator
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	user, err := userservice.New(userservice.WithGormPostgeReposiotry(db))
	if err != nil {
		panic(err)
	}

	return &UserHandler{
		user:      user,
		validator: utils.NewValidator(),
	}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.user.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "GET all users success",
		Data:    users,
	})
}

func (h *UserHandler) GetByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := h.user.GetByUsername(username)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "GET by username success",
		Data:    user,
	})
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	req := userservice.CreateUser{}
	isNotValid, err := utils.ParseAndValidateRequest(c, &req, h.validator)
	if isNotValid || err != nil {
		return err
	}

	user, err := h.user.CreateUser(req)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "CREATE user success",
		Data:    user,
	})
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	username := c.Params("username")
	req := userservice.UpdateUser{}
	isNotValid, err := utils.ParseAndValidateRequest(c, &req, h.validator)
	if isNotValid || err != nil {
		return err
	}

	user, err := h.user.UpdateUser(username, req)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "UPDATE user success",
		Data:    user,
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	username := c.Params("username")

	err := h.user.DeleteUser(username)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "DELETE profile success",
	})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	authUser := c.Locals("user").(*userservice.User)
	acc, err := h.user.GetPofile(*authUser)
	if err != nil {
		return err
	}

	data, err := utils.ProfileResponse(acc)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "GET profile success",
		Data:    data,
	})
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	authUser := c.Locals("user").(*userservice.User)
	var acc account.Account

	if *authUser.RoleID == 1 {
		return c.Status(fiber.StatusBadRequest).JSON(types.Response{
			Error:   true,
			Message: "Admin can't update profile",
		})
	} else if *authUser.RoleID == 2 {
		req := userservice.ProfileStudent{}
		isNotValid, err := utils.ParseAndValidateRequest(c, &req, h.validator)
		if isNotValid || err != nil {
			return err
		}

		acc, err = h.user.UpdateProfileStudent(*authUser, req)
		if err != nil {
			return err
		}
	} else if *authUser.RoleID == 3 {
		req := userservice.ProfileLecturer{}
		isNotValid, err := utils.ParseAndValidateRequest(c, &req, h.validator)
		if isNotValid || err != nil {
			return err
		}

		acc, err = h.user.UpdateProfileLecturer(*authUser, req)
		if err != nil {
			return err
		}
	} else {
		return errors.New("internal server error")
	}

	data, err := utils.ProfileResponse(acc)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "UPDATE profile success",
		Data:    data,
	})
}

func (h *UserHandler) Authorized(c *fiber.Ctx) error {
	req := userservice.UserAuth{}
	isNotValid, err := utils.ParseAndValidateRequest(c, &req, h.validator)
	if isNotValid || err != nil {
		return err
	}

	auth, err := h.user.AuthAPI(req)
	if err != nil {
		return err
	}

	return c.JSON(types.Response{
		Error:   false,
		Message: "Authorized",
		Data:    auth,
	})
}
