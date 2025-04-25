package handlers

import (
	"backend-server/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) CreateRequestEntry(ctx *fiber.Ctx) error {
	requestBody := &models.CreateRequestBodyRequest{}
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return &fiber.Error{
			Code:    fiber.StatusBadGateway,
			Message: "error while parsing the requestBody: " + err.Error(),
		}
	}
	if requestBody.Name == "" || requestBody.Promo_Text == "" {
		return &fiber.Error{
			Code:    404,
			Message: "Please check requestbody either name or promo_text is empty",
		}
	}
	userId := ctx.Query("user-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "userId in the query parameter are missing",
		})
	}
	log.Println("the userid = ", userId)
	resp, err := h.ReqService.CreateRequestEntry(requestBody, uuid.MustParse(userId))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.Status(200).JSON(models.ServiceResponse{
		Code: 200,
		Message: resp.Message,
		Data: nil,
	})
}

func (h *Handler) ListAllRequestEntry(ctx *fiber.Ctx) error {
	userId := ctx.Query("user-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "userId in the query parameter are missing",
		})
	}
	resp, err := h.ReqService.ListAllRequestEntry(uuid.MustParse(userId))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.Status(200).JSON(models.ServiceResponse{
		Code: 200,
		Message: "The request for this user",
		Data: resp,
	})
}

func (h *Handler) FindRequestEntry(ctx *fiber.Ctx) error {
	userId := ctx.Query("user-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "userId in the query parameter are missing",
		})
	}
	reqId := ctx.Query("req-id")
	if userId == "" {
		return ctx.Status(502).JSON(fiber.Map{
			"message": "requestBody Id in the query parameter are missing",
		})
	}
	resp, err := h.ReqService.FindRequestEntry(uuid.MustParse(reqId), uuid.MustParse(userId))
	if err != nil {
		if serviceErr, ok := err.(*models.ServiceResponse); ok {
			return ctx.Status(serviceErr.Code).JSON(err)
		} else {
			return ctx.JSON(500, "an unexpected error occurred")
		}
	}
	return ctx.Status(200).JSON(models.ServiceResponse{
		Code: 200,
		Message: "the specific promo text info",
		Data: resp,
	})
}