package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateJobEntry(ctx *fiber.Ctx) error {
	var requestBody = &models.CreateJobRequest{}
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}

	if requestBody.Name == "" || requestBody.Type == "" {
		log.Println("the requestBody: ", requestBody)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "either name or type is missing in the request Body",
		}
	}

	userId := ctx.Query("user-id")
	fileId := ctx.Query("file-id")
	resp, err := h.JobService.CreateJobEntry(requestBody, uuid.MustParse(userId), uuid.MustParse(fileId))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: resp.Message,
	})
}

func (h *Handler) ListJobEntry(ctx *fiber.Ctx) error {
	id := ctx.Query("user-id")
	if id == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "userId is required",
		})
	}
	resp, err := h.JobService.ListJobEntry(uuid.MustParse(id))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: "",
		Data:    resp,
	})
}
