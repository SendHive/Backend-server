package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New()
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("error while starting the server: ", err)
	}

}
