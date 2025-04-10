package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateUserEntry(ctx *fiber.Ctx) error {
	var requestBody = &models.CreateUserRequest{}
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
	if requestBody.Name == "" || requestBody.Email == "" || requestBody.Password == "" {
		return &fiber.Error{
			Code:    404,
			Message: "Name, Email or Pasword  in the requestBody is the required field",
		}
	}
	resp, err := h.UserService.CreateUserEntry(requestBody)
	if err != nil {
		log.Println("error while creating the user: " + err.Error())
		return ctx.JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: resp.Message,
	})
}

func (h *Handler) GetUserQRCodeImage(ctx *fiber.Ctx) error {
	id := ctx.Query("user-id")
	if id == "" {
		return &models.ServiceResponse{
			Code:    500,
			Message: "user-id in the query is empty",
		}
	}
	resp, err := h.UserService.GetUserQRCodeImage(uuid.MustParse(id))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	ctx.Set("Content-Type", "image/png")
	return ctx.Send([]byte(resp))
}

func (h *Handler) UserAuthentication(ctx *fiber.Ctx) error {
	var requestBody = &models.UserAuthenticationRequest{}
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}
	userId := ctx.Query("id")
	resp, err := h.UserService.UserAuthentication(requestBody, uuid.MustParse(userId))
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
