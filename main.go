package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/consumer"
	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/mongodb"
	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/rabbitmq"
	"github.com/imonasterio/go-mongodb-rabbitmq/cmd/api/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("Starting the application...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client := mongodb.ConnectMongo()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	connectRabbitmq := rabbitmq.ConnectRabbit()

	defer connectRabbitmq.Close()

	ch := consumer.StartConsumer(connectRabbitmq)

	defer ch.Close()

	server.InitServer()
}
