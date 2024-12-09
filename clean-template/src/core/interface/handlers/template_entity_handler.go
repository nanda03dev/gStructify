package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/application/services"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type TemplateEntityHandler interface {
	CreateTemplateEntity(ctx *fiber.Ctx) error
	GetTemplateEntityByID(ctx *fiber.Ctx) error
	UpdateTemplateEntityById(ctx *fiber.Ctx) error
	DeleteTemplateEntityById(ctx *fiber.Ctx) error
}

type templateEntityHandler struct {
	templateEntityService services.TemplateEntityService
}

func NewTemplateEntityHandler(templateEntityService services.TemplateEntityService) TemplateEntityHandler {
	return &templateEntityHandler{
		templateEntityService: templateEntityService,
	}
}

func (c *templateEntityHandler) CreateTemplateEntity(ctx *fiber.Ctx) error {
	var templateEntityDTO dto.CreateTemplateEntityDTO

	if err := ctx.BodyParser(&templateEntityDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": common.InvalidRequestError})
	}

	result, err := c.templateEntityService.Create(templateEntityDTO)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": common.InvalidRequestError})
	}

	return ctx.Status(http.StatusOK).JSON(result)
}

func (c *templateEntityHandler) GetTemplateEntityByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	templateEntity, err := c.templateEntityService.GetById(idParam)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": common.TemplateEntityNotFoundError})
	}

	return ctx.Status(http.StatusOK).JSON(templateEntity)
}

func (c *templateEntityHandler) UpdateTemplateEntityById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	_, err := c.templateEntityService.GetById(idParam)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": common.TemplateEntityNotFoundError})
	}

	var templateEntityDTO dto.UpdateTemplateEntityDTO

	if err := ctx.BodyParser(&templateEntityDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": common.InvalidRequestError})
	}

	result, err := c.templateEntityService.Update(idParam, templateEntityDTO)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": common.InvalidRequestError})
	}

	return ctx.Status(http.StatusOK).JSON(result)
}

func (c *templateEntityHandler) DeleteTemplateEntityById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	_, err := c.templateEntityService.GetById(idParam)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": common.TemplateEntityNotFoundError})
	}

	deleteErr := c.templateEntityService.Delete(idParam)

	if deleteErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": common.ErrorDeletingData})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"data": common.DataDeletedSuccessfully})
}
