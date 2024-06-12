package common

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func ConnectRabbitMQ(user, password, host, port, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rc RabbitClient) CreateExchange(name string, kind string) error {
	err := rc.ch.ExchangeDeclare(name, kind, true, false, false, false, amqp.Table{})
	if err != nil {
		log.Fatalf("Error while creatint exchange %s", name)
	}
	return err

}

func (rc RabbitClient) Close() error {
	return rc.ch.Close()
}
