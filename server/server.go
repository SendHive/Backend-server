package server

import (
	"backend-server/external"
	"backend-server/handlers"
	"backend-server/services"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New()

	//Setting up the database
	err := external.ConnectDB()
	if err != nil {
		return
	}

	ser, err := services.NewSmtpServiceRequest()
	if err != nil {
		log.Println("error while starting the smtp services: ", err)
		return
	}

	handler := handlers.NewHandler().SmtpHandler(ser)
	Routes(app, handler)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("error while starting the server: ", err)
	}

}
