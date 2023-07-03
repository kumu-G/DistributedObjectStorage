package rabbitmq

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	channel  *amqp091.Channel
	Name     string
	exchange string
}

func New(s string) *RabbitMQ {
	conn, e := amqp091.Dial(s)
	if e != nil {
		panic(e)
	}
	ch, e := conn.Channel()
	if e != nil {
		panic(e)
	}
	q, e := ch.QueueDeclare(
		"",    // Name
		false, //durable
		true,  //delete when unused
		false, //exclusive
		false, //  no-wait
		nil,   // arguments
	)
	if e != nil {
		panic(e)
	}
	mq := new(RabbitMQ)
	mq.channel = ch
	mq.Name = q.Name
	return mq
}

func (q *RabbitMQ) Bind(exchange string) {
	e := q.channel.QueueBind(
		q.Name,
		"",
		exchange,
		false,
		nil)
	if e != nil {
		panic(e)
	}
	q.exchange = exchange
}

func (q *RabbitMQ) Send(queue string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp091.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}
}

func (q *RabbitMQ) Publish(exchange string, body interface{}) {
	str, e := json.Marshal(body)
	if e != nil {
		panic(e)
	}
	e = q.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp091.Publishing{
			ReplyTo: q.Name,
			Body:    []byte(str),
		})
	if e != nil {
		panic(e)
	}
}
func (q *RabbitMQ) Consume() <-chan amqp091.Delivery {
	c, e := q.channel.Consume(q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if e != nil {
		panic(e)
	}
	return c
}

func (q *RabbitMQ) Close() {
	q.channel.Close()
}
