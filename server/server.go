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

	// Setting the external queue connection
	qConn, Iq, err := external.SetupQueue()

	if err != nil {
		log.Println(err)
		return
	}

	qu, err := external.DeclareQueue(qConn, Iq)
	if err != nil {
		log.Println(err)
		return
	}

	//Seting connection to the minio
	mc, mI, err := external.ConnectMinio()
	if err != nil {
		log.Println(err)
		return
	}

	Jser, err := services.NewJobServiceRequest(qu, qConn, mc, mI)
	if err != nil {
		log.Println("error while setting the new job service: ", err)
		return
	}

	ser, err := services.NewSmtpServiceRequest()
	if err != nil {
		log.Println("error while starting the smtp services: ", err)
		return
	}

	user, err := services.NewUserServiceReqest()
	if err != nil {
		log.Println("error while starting the user service: ", err)
		return
	}

	login, err := services.NewLoginService()
	if err != nil {
		log.Println("error while starting the login service: ", err)
		return
	}

	handler := handlers.NewHandler().SmtpHandler(ser).JobHandler(Jser).UserHandler(user).LoginHandler(login)
	Routes(app, handler)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal("error while starting the server: ", err)
	}

}
