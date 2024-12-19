package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nanda03dev/go-ms-template/src/common"
	"github.com/nanda03dev/go-ms-template/src/core/application/services"
	"github.com/nanda03dev/go-ms-template/src/core/domain/aggregates"
	"github.com/nanda03dev/go-ms-template/src/core/interface/dto"
)

type TemplateEntityHandler interface {
	CreateTemplateEntity(ctx *fiber.Ctx) error
	GetTemplateEntityByID(ctx *fiber.Ctx) error
	FindTemplateEntityWithFilter(ctx *fiber.Ctx) error
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
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.InvalidRequestError))
	}

	result, err := c.templateEntityService.Create(templateEntityDTO)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.InvalidRequestError))
	}

	return ctx.Status(http.StatusOK).JSON(SuccessResponse(c.toResponseDTO(result)))
}

func (c *templateEntityHandler) GetTemplateEntityByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	templateEntity, err := c.templateEntityService.GetById(idParam)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(ErrorResponse(common.TemplateEntityNotFoundError))
	}

	return ctx.Status(http.StatusOK).JSON(SuccessResponse(c.toResponseDTO(templateEntity)))
}

func (c *templateEntityHandler) FindTemplateEntityWithFilter(ctx *fiber.Ctx) error {
	var filterDTO common.FilterQueryDTO

	if err := ctx.BodyParser(&filterDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.InvalidRequestError))
	}

	templateEntitys, err := c.templateEntityService.FindWithFilter(filterDTO)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(ErrorResponse(common.TemplateEntityNotFoundError))
	}

	return ctx.Status(http.StatusOK).JSON(SuccessResponse(c.toResponseDTOArray(templateEntitys)))
}

func (c *templateEntityHandler) UpdateTemplateEntityById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	_, err := c.templateEntityService.GetById(idParam)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(ErrorResponse(common.TemplateEntityNotFoundError))
	}

	var templateEntityDTO dto.UpdateTemplateEntityDTO

	if err := ctx.BodyParser(&templateEntityDTO); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.InvalidRequestError))
	}

	result, err := c.templateEntityService.Update(idParam, templateEntityDTO)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.InvalidRequestError))
	}

	return ctx.Status(http.StatusOK).JSON(SuccessResponse(c.toResponseDTO(result)))
}

func (c *templateEntityHandler) DeleteTemplateEntityById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")

	_, err := c.templateEntityService.GetById(idParam)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(ErrorResponse(common.TemplateEntityNotFoundError))
	}

	deleteErr := c.templateEntityService.Delete(idParam)

	if deleteErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse(common.ErrorDeletingData))
	}

	return ctx.Status(http.StatusOK).JSON(SuccessResponse(common.DataDeletedSuccessfully))
}

// Helper function to convert Entity to TemplateEntityResponseDTO
func (c *templateEntityHandler) toResponseDTO(templateEntity *aggregates.TemplateEntity) dto.TemplateEntityResponseDTO {
	return dto.TemplateEntityResponseDTO{
		ID: templateEntity.ID,
	}
}

// Function to convert an array of TemplateEntitys to an array of TemplateEntityResponseDTOs
func (c *templateEntityHandler) toResponseDTOArray(templateEntitys []*aggregates.TemplateEntity) []dto.TemplateEntityResponseDTO {
	var responseDTOs = make([]dto.BookResponseDTO, 0, len(templateEntitys))
	for _, templateEntity := range templateEntitys {
		responseDTOs = append(responseDTOs, c.toResponseDTO(templateEntity))
	}
	return responseDTOs
}
