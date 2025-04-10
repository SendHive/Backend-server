package handlers

import (
	"backend-server/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateFileEntry(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).SendString("File is required")
	}
	userId := ctx.Query("user-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "userId in the query parameter are missing",
		})
	}
	resp, err := h.FileService.CreateFileEntry(&models.CreateFileRequest{}, file, uuid.MustParse(userId))
	if err != nil {
		return ctx.JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: resp.Message,
	})
}

func (h *Handler) ListFiles(ctx *fiber.Ctx) error {
	userId := ctx.Query("user-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "userId in the query parameter are missing",
		})
	}
	resp, err := h.FileService.ListFiles(uuid.MustParse(userId))
	if err != nil {
		return ctx.JSON(err)
	}
	if len((resp)) == 0 {
		return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
			Code:    200,
			Message: "No file for the current user",
			Data:    resp,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: "The files as per the user details",
		Data:    resp,
	})
}
