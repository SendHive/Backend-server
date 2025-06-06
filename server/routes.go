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
	job.Get("/", h.ListJobEntry)

	user := app.Group("/user")
	user.Post("/", h.CreateUserEntry)
	user.Get("/qr", h.GetUserQRCodeImage)
	user.Post("/auth", h.UserAuthentication)

	login := app.Group("/login")
	login.Post("/", h.CreateLoginEntry)

	file := app.Group("/file")
	file.Post("/", h.CreateFileEntry)
	file.Get("/", h.ListFiles)

	req := app.Group("/req")
	req.Post("/", h.CreateRequestEntry)
	req.Get("/", h.ListAllRequestEntry)
	req.Get("/find", h.FindRequestEntry)
	req.Put("/", h.UpdateRequestEntry)

}
