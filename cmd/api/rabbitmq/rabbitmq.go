package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbit() *amqp.Connection {

	//run with Docker
	//docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
	//user; guest, password: guest
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	return conn

}
