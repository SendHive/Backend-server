package handlers

import (
	"backend-server/models"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateSmtpEntry(ctx *fiber.Ctx) error {
	var requestBody *models.CreateSmtpEntryRequest
	if h == nil {
		return &models.ServiceResponse{
			Message: "Handler is not intialized",
		}
	}
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}
	if requestBody.Port == "" || requestBody.Server == "" || requestBody.Password == "" || requestBody.Username == "" {
		return &fiber.Error{
			Code:    404,
			Message: "Please check requestbody either port, server , username or password is empty",
		}
	}

	resp, err := h.SmtpService.CreateSmtpEntry(requestBody)
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

func (h *Handler) UpdateSmtpEntry(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
	var requestBody *models.UpdateSmtpEntryRequest
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}
	if requestBody.Port == "" || requestBody.Server == "" || requestBody.Password == "" || requestBody.Username == "" {
		return &fiber.Error{
			Code:    404,
			Message: "Please check requestbody either port, server , username or password is empty",
		}
	}
	resp, err := h.SmtpService.UpdateSmtpEntry(id, requestBody)
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

func (h *Handler) ListSmtpEntry(ctx *fiber.Ctx) error {
	id := ctx.Query("user-id")
	if id == "" {
		return &fiber.Error{
			Code:    404,
			Message: "user-id not found",
		}
	}
	resp, err := h.SmtpService.ListSmtpEntry(id)
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
