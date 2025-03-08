package external

import (
	"backend-server/models"
	"log"
	"time"

	queue "github.com/SendHive/Infra-Common/queue"
	"github.com/rabbitmq/amqp091-go"
)

func SetupQueue() (*amqp091.Connection, queue.IQueueService, error) {
	qConn, err := queue.NewQueueRequest()
	if err != nil {
		log.Fatal("the error while creating the queue instance: ", err)
		return nil, nil, err
	}
	time.Sleep(3 * time.Second)
	qconn, err := qConn.Connect()
	if err != nil {
		return nil, nil, err
	}
	time.Sleep(3 * time.Second)
	return qconn, qConn, nil
}

func DeclareQueue(qConn *amqp091.Connection, Iq queue.IQueueService) (qu amqp091.Queue, err error) {
	queue, err := Iq.DeclareQueue(qConn)

	if err != nil {
		return amqp091.Queue{}, err
	}
	return queue, nil

}

func PublishMessage(q amqp091.Queue, Iq queue.IQueueService, conn *amqp091.Connection, body string) error {
	err := Iq.PublishMessage(q, conn, body)
	if err != nil {
		return &models.ServiceResponse{
			Code:    500,
			Message: err.Error(),
		}
	}
	return &models.ServiceResponse{
		Message: "Published message successfully",
	}
}
