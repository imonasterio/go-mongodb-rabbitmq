package producer

import (
	"fmt"
	"log"
	"time"

	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/models"
	"github.com/streadway/amqp"
)

func StartProducer(conn *amqp.Connection, tweets models.Tweet) {

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare("gophers", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	//debug only
	fmt.Println(q)

	for msg := range tweets.User.ScreenName {
		err := ch.Publish("", q.Name, false, false,
			amqp.Publishing{
				Headers:     nil,
				ContentType: "text/plain",
				Body:        []byte(fmt.Sprintf("%v", msg)),
			})

		if err != nil {
			break
		}

		//wait 2 seconds until send another message
		time.Sleep(2 * time.Second)
	}
}
