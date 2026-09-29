package http

import (
	"go-postgresql-starter-kit/internal/delivery/http/middleware"
	"go-postgresql-starter-kit/internal/model"
	"go-postgresql-starter-kit/internal/usecase"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ContactController struct {
	UseCase *usecase.ContactUseCase
	Log     *logrus.Logger
}

func NewContactController(useCase *usecase.ContactUseCase, log *logrus.Logger) *ContactController {
	return &ContactController{
		UseCase: useCase,
		Log:     log,
	}
}

// Create godoc
// @Summary      Create contact
// @Description  Create a new contact for the authenticated user
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        request body model.CreateContactRequest true "Contact details"
// @Success      200 {object} model.WebResponse[model.ContactResponse]
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/contacts [post]
func (c *ContactController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.CreateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}
	request.UserId = auth.ID

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// List godoc
// @Summary      List contacts
// @Description  Search and paginate contacts of the authenticated user
// @Tags         Contacts
// @Produce      json
// @Security     ApiKeyAuth
// @Param        name  query string false "Contact name filter"
// @Param        email query string false "Contact email filter"
// @Param        phone query string false "Contact phone filter"
// @Param        page  query int    false "Page number" default(1)
// @Param        size  query int    false "Page size" default(10)
// @Success      200 {object} model.PageResponse[model.ContactResponse]
// @Failure      401 {object} model.ErrorResponse
// @Router       /api/contacts [get]
func (c *ContactController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.SearchContactRequest{
		UserId: auth.ID,
		Name:   ctx.Query("name", ""),
		Email:  ctx.Query("email", ""),
		Phone:  ctx.Query("phone", ""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.PageResponse[model.ContactResponse]{
		Data:   responses,
		Paging: paging,
	})
}

// Get godoc
// @Summary      Get contact by ID
// @Description  Retrieve single contact by its ID
// @Tags         Contacts
// @Produce      json
// @Security     ApiKeyAuth
// @Param        contactId path string true "Contact ID"
// @Success      200 {object} model.WebResponse[model.ContactResponse]
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/contacts/{contactId} [get]
func (c *ContactController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &model.GetContactRequest{
		UserId: auth.ID,
		ID:     ctx.Params("contactId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// Update godoc
// @Summary      Update contact
// @Description  Update an existing contact by ID
// @Tags         Contacts
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        contactId path string true "Contact ID"
// @Param        request   body model.UpdateContactRequest true "Updated contact details"
// @Success      200 {object} model.WebResponse[model.ContactResponse]
// @Failure      400 {object} model.ErrorResponse
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/contacts/{contactId} [put]
func (c *ContactController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(model.UpdateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.UserId = auth.ID
	request.ID = ctx.Params("contactId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.ContactResponse]{Data: response})
}

// Delete godoc
// @Summary      Delete contact
// @Description  Delete a contact by ID
// @Tags         Contacts
// @Produce      json
// @Security     ApiKeyAuth
// @Param        contactId path string true "Contact ID"
// @Success      200 {object} model.WebResponse[bool]
// @Failure      401 {object} model.ErrorResponse
// @Failure      404 {object} model.ErrorResponse
// @Router       /api/contacts/{contactId} [delete]
func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &model.DeleteContactRequest{
		UserId: auth.ID,
		ID:     contactId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
