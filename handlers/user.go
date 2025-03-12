package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
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
	if requestBody.Name == "" {
		return &fiber.Error{
			Code:    404,
			Message: "Name in the requestBody is the required field",
		}
	}
	resp, err := h.UserService.CreateUserEntry(requestBody)
	if err != nil {
		log.Println("error : " + err.Error())
		return ctx.JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(models.ServiceResponse{
		Code:    200,
		Message: resp.Message,
	})
}
