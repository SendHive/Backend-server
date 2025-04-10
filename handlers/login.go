package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateLoginEntry(ctx *fiber.Ctx) error {
	var requestBody *models.CreateLoginRequest
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}
	if requestBody.Email == ""  {
		log.Println("Email = ", requestBody.Email)
		return &models.ServiceResponse{
			Code:    404,
			Message: "Either name or secretCode is empty in the requestBody",
		}
	}
	resp, err := h.LoginService.CreateLoginEntry(requestBody)
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.JSON(resp)

}
