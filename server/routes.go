package server

import (
	"backend-server/handlers"

	fiber "github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	app.Get("/check", func(c *fiber.Ctx) error {
		return c.SendString("server api's are healthy")
	})

	smtp := app.Group("/smtp")
	smtp.Get("/", h.ListSmtpEntry)
	smtp.Post("/", h.CreateSmtpEntry)
	smtp.Put("/", h.UpdateSmtpEntry)

	job := app.Group("/job")
	job.Post("/", h.CreateJobEntry)
}
