package services

import (
	"log"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
)

type MessageQueue interface {
	Publisher() error
	Consumer() error
}

// Publisher publish event to queue
type Publisher struct {
	Uri          string
	QueueName    string
	Exchange     string
	ExchangeType string
	Body         string
	Reliable     bool
}

func (p Publisher) Publish() {
	log.Println("[-] Connecting to", p.Uri)
	connection, err := connect(p.Uri)

	if err != nil {
		log.Fatalf("[x] AMQP connection error: %s", err)
	}

	log.Println("[√] Connected successfully")

	channel, err := connection.Channel()

	if err != nil {
		log.Fatalf("[x] Failed to open a channel: %s", err)
	}

	defer channel.Close()

	log.Println("[-] Declaring Exchange", p.ExchangeType, p.Exchange)
	err = channel.ExchangeDeclare(p.Exchange, p.ExchangeType, nil)

	if err != nil {
		log.Fatalf("[x] Failed to declare exchange: %s", err)
	}
	log.Println("[√] Exchange", p.Exchange, "has been declared successfully")

	log.Println("[-] Declaring queue", p.QueueName, "into channel")
	queue, err := declareQueue(p.QueueName, channel)

	if err != nil {
		log.Fatalf("[x] Queue could not be declared. Error: %s", err.Error())
	}
	log.Println("[√] Queue", p.QueueName, "has been declared successfully")

	if p.Reliable {
		log.Printf("[-] Enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			log.Fatalf("[x] Channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan wabbit.Confirmation, 1))

		defer confirmOne(confirms)
	}

	log.Println("[-] Sending message to queue:", p.QueueName, "- p.exchange:", p.Exchange)
	log.Println("\t", p.Body)

	err = publishMessage(p.Body, p.Exchange, queue, channel)

	if err != nil {
		log.Fatalf("[x] Failed to publish a message. Error: %s", err.Error())
	}
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

func publishMessage(body string, exchange string, queue wabbit.Queue, channel wabbit.Channel) error {
	return channel.Publish(
		exchange,     // exchange
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
