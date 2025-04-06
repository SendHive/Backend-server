package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateJobEntry(ctx *fiber.Ctx) error {
	var requestBody = &models.CreateJobRequest{}
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	} else {
		requestBody.Name = name
	}

	log.Println("the name : ", name)

	userId := ctx.Query("file-id")
	fileId := ctx.Query("file-id")
	log.Println("the userid:", userId)
	resp, err := h.JobService.CreateJobEntry(requestBody, uuid.MustParse(userId), uuid.MustParse(fileId))
	if err != nil {
		return ctx.JSON(err)
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
		return ctx.JSON(err)
	}
	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: "",
		Data:    resp,
	})
}
