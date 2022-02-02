package rabbitMQ

import (
	"log"

	"github.com/NeowayLabs/wabbit"
	"github.com/NeowayLabs/wabbit/amqp"
)

type AmqpChannel interface {
	//ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool) error
	//QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error)
	//QueueBind(name, key, exchange string, noWait bool) error
	//Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error)
	Publish(uri string, queueName string, exchange string, exchangeType string, body string, reliable bool) error
}

type amqpService struct{}

func NewAmqpChannel() AmqpChannel {
	return &amqpService{}
}

func (a amqpService) Publish(uri string, queueName string, exchange string, exchangeType string, body string, reliable bool) error {
	log.Println("[-] Connecting to", uri)
	connection, err := connect(uri)

	if err != nil {
		log.Fatalf("[x] AMQP connection error: %s", err)
	}

	log.Println("[√] Connected successfully")

	channel, err := connection.Channel()

	if err != nil {
		log.Fatalf("[x] Failed to open a channel: %s", err)
	}

	defer channel.Close()

	log.Println("[-] Declaring Exchange", exchangeType, exchange)
	err = channel.ExchangeDeclare(exchange, exchangeType, nil)

	if err != nil {
		log.Fatalf("[x] Failed to declare exchange: %s", err)
	}
	log.Println("[√] Exchange", exchange, "has been declared successfully")

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

	log.Println("[-] Sending message to queue:", queueName, "- exchange:", exchange)
	log.Println("\t", body)

	err = publishMessage(body, exchange, queue, channel)

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

//type Publisher interface {
//	PublishEvent() error
//}
//type publisherService struct{}
//
//func NewPublisherService() Publisher {
//	return &publisherService{}
//}
//
//func (p publisherService) PublishEvent() error {
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//	failOnError(err, "Failed to connect to rabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	err = ch.ExchangeDeclare(
//		"logs_topic", // name
//		"topic",      // type
//		true,         // durable
//		false,        // auto-deleted
//		false,        // internal
//		false,        // no-wait
//		nil,          // arguments
//	)
//	failOnError(err, "Failed to declare an exchange")
//
//	body := bodyFrom(os.Args)
//	err = ch.Publish(
//		"logs_topic",          // exchange
//		severityFrom(os.Args), // routing key
//		false,                 // mandatory
//		false,                 // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//		})
//	failOnError(err, "Failed to publish a message")
//
//	log.Printf(" [x] Sent %s", body)
//	return nil
//}
//
//func bodyFrom(args []string) string {
//	var s string
//	if (len(args) < 3) || os.Args[2] == "" {
//		s = "hello"
//	} else {
//		s = strings.Join(args[2:], " ")
//	}
//	return s
//}
//
//func severityFrom(args []string) string {
//	var s string
//	if (len(args) < 2) || os.Args[1] == "" {
//		s = "anonymous.info"
//	} else {
//		s = os.Args[1]
//	}
//	return s
//}
//
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}
