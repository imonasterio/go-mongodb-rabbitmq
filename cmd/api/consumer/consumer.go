package consumer

import (
	"fmt"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func StartConsumer(conn *amqp.Connection) *amqp.Channel {

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	chDelivery, err := ch.Consume(
		"gophers",
		"",
		true,
		false,
		false,
		false, nil)

	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for delivery := range chDelivery {
			msg := string(delivery.Body)
			fmt.Println("msg: " + strings.ToUpper(msg))
		}
	}()

	return ch
}
