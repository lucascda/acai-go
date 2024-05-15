package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucas_cda/acai-go/common"
)

func main() {
	godotenv.Load()
	rabbitmq := struct {
		user     string
		password string
		host     string
		port     string
		vhost    string
	}{
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
		os.Getenv("RABBITMQ_VHOST"),
	}

	conn, err := common.ConnectRabbitMQ(rabbitmq.user, rabbitmq.password, rabbitmq.host, rabbitmq.port, rabbitmq.vhost)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client, err := common.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	log.Println(client)

}
