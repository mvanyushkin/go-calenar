package messages

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type QueueClient struct {
	connectionString string
	conn             *amqp.Connection
	queue            amqp.Queue
	ch               *amqp.Channel
}

func NewQueueClient(connectionString string) (*QueueClient, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare("calendar_queue", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = ch.Qos(1, 0, false)
	if err != nil {
		return nil, err
	}

	return &QueueClient{
		connectionString: connectionString,
		conn:             conn,
		ch:               ch,
		queue:            queue,
	}, nil
}

func (q QueueClient) Close() error {
	err := q.ch.Close()
	if err != nil {
		return err
	}

	return q.conn.Close()
}

func (q QueueClient) SendMessage(message MessageDto) error {
	marshalled, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = q.ch.Publish(
		"",
		q.queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         marshalled,
		})

	return err
}

func (q QueueClient) ReceiveMessages() (<-chan WorkItem, error) {
	messageQueue, err := q.ch.Consume(q.queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	outputChannel := make(chan WorkItem)

	go func() {
		for d := range messageQueue {
			log.Printf("Received a Message.")
			m := MessageDto{}
			err := json.Unmarshal(d.Body, &m)
			if err != nil {
				d.Reject(false)
			}

			workItem := WorkItem{
				Message: m,
				DoneCallback: func() {
					d.Ack(false)
				},
			}

			outputChannel <- workItem
		}
	}()

	return outputChannel, nil
}
