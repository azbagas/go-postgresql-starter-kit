package http

import (
	"github.com/azbagas/go-postgresql-starter-kit/internal/delivery/http/middleware"
	"github.com/azbagas/go-postgresql-starter-kit/internal/model"
	"github.com/azbagas/go-postgresql-starter-kit/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Register a new user in the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body model.RegisterUserRequest true "User Registration Info"
// @Success      200 {object} model.WebResponse[model.UserResponse]
// @Failure      400 {object} model.ErrorResponse
// @Router       /api/users [post]
func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and obtain authentication token
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        request body model.LoginUserRequest true "User Login Info"
// @Success      200 {object} model.WebResponse[model.UserResponse]
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/users/_login [post]
func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(model.LoginUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Current godoc
// @Summary      Get current user profile
// @Description  Retrieve profile information for the authenticated user
// @Tags         Users
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} model.WebResponse[model.UserResponse]
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/users/_current [get]
func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Current(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to get current user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}

// Logout godoc
// @Summary      Logout current user
// @Description  Invalidate authentication token for the current user
// @Tags         Users
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} model.WebResponse[bool]
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/users [delete]
func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.LogoutUserRequest{
		ID: auth.ID,
	}

	response, err := c.UseCase.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}

// Update godoc
// @Summary      Update current user
// @Description  Update profile of current authenticated user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request body model.UpdateUserRequest true "User Update Info"
// @Success      200 {object} model.WebResponse[model.UserResponse]
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/users/_current [patch]
func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.ID = auth.ID
	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.UserResponse]{Data: response})
}
