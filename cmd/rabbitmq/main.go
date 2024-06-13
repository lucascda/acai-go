package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucas_cda/go-todo-microservices/common"
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
	log.Print("Succesfully connect to RabbitMQ")

	client, err := common.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	log.Print("Succesfully RabbitMQ client")

	err = client.CreateExchange("auth_events", "direct")
	if err != nil {
		panic(err)
	}
	log.Print("Succesfully created exchange")

	if err = client.CreateQueue("auth_signup"); err != nil {
		panic(err)
	}
	log.Print("Succesfully created queue")

	if err = client.CreateBinding("auth_signup", "auth.signup", "auth_events"); err != nil {
		panic(err)
	}
	log.Print("Succesfully created binding")
}
