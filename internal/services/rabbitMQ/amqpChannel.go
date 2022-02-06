package rabbitMQ

import (
	"log"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
)

type AmqpChannel interface {
	Channel() (wabbit.Channel, error)
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error)
	Publish(queueName string, exchange string, exchangeType string, body string, reliable bool) error
}

type AmqpService struct {
	Uri string
}

func NewAmqpChannel(a AmqpService) AmqpChannel {
	return &a
}

func (a AmqpService) Channel() (wabbit.Channel, error) {
	log.Println("[-] Connecting to", a.Uri)
	connection, err := connect(a.Uri)

	if err != nil {
		log.Fatalf("[x] AMQP connection error: %s", err)
	}

	log.Println("[√] Connected successfully")

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("[x] Failed to open a channel: %s", err)
	}

	return channel, nil
}

func (a AmqpService) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
	//TODO implement me
	panic("implement me")
}

func (a AmqpService) Publish(queueName string, exchange string, exchangeType string, body string, reliable bool) error {
	channel, err := a.Channel()

	if err != nil {
		log.Fatalf("[x] Failed to open a channel: %s", err)
	}

	defer channel.Close()

	log.Println("[-] Declaring queue", queueName, "into channel")
	queue, err := declareQueue(queueName, channel)
	if err != nil {
		log.Fatalf("[x] Queue could not be declared. Error: %s", err.Error())
	}

	log.Println("[√] Queue", queueName, "has been declared successfully")

	if reliable {
		log.Printf("[-] Enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			log.Fatalf("[x] Channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan wabbit.Confirmation, 1))

		defer confirmOne(confirms)
	}

	log.Println("[-] Sending message to", queueName)
	log.Println("\t", body)

	err = publishMessage(body, queue, channel)

	if err != nil {
		log.Fatalf("[x] Failed to publish a message. Error: %s", err.Error())
	}
	return nil
}

func connect(uri string) (*amqp.Conn, error) {
	return amqp.Dial(uri)
}

func declareQueue(queueName string, channel wabbit.Channel) (wabbit.Queue, error) {
	return channel.QueueDeclare(
		queueName,
		wabbit.Option{
			"durable":    true,
			"autoDelete": false,
			"exclusive":  false,
			"noWait":     false,
		},
	)
}

func publishMessage(body string, queue wabbit.Queue, channel wabbit.Channel) error {
	return channel.Publish(
		"",           // exchange
		queue.Name(), // routing key
		[]byte(body),
		wabbit.Option{
			"deliveryMode": 2,
			"contentType":  "text/plain",
		})
}

func confirmOne(confirms <-chan wabbit.Confirmation) {
	log.Printf("[-] Waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack() {
		log.Printf("[√] Confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("[x] Failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
