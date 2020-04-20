package messages

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"sync/atomic"
	"time"
)

type QueueClient struct {
	connectionString        string
	conn                    *amqp.Connection
	queue                   amqp.Queue
	ch                      *amqp.Channel
	closed                  int32
	connectionEstablishedCh chan bool
	connectionClosedCh      chan bool
}

func NewQueueClient(connectionString string) (*QueueClient, error) {
	c := &QueueClient{
		connectionString:        connectionString,
		connectionEstablishedCh: make(chan bool),
	}

	c.ensureConnected()
	<-c.connectionEstablishedCh
	return c, nil
}

func (q *QueueClient) ensureConnected() {
	go func() {
		for {
			err := q.establishConnection()
			if err == nil {
				q.connectionClosedCh = make(chan bool)
				q.setIsClosed(false)
				q.connectionEstablishedCh <- true
				reason, _ := <-q.conn.NotifyClose(make(chan *amqp.Error))
				q.setIsClosed(true)
				q.connectionClosedCh <- true
				log.Errorf("connection to rabbitmq is lost, reason: %v", reason)
			} else {
				time.Sleep(time.Second)
			}
		}
	}()
}

func (q *QueueClient) establishConnection() error {
	log.Infof("dialing to %v ...", q.connectionString)
	conn, err := amqp.Dial(q.connectionString)
	if err != nil {
		log.Errorf("an error has occurred %v", err.Error())
		return err
	}
	log.Info("done.")

	log.Info("creating channel...")
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	log.Info("done.")

	log.Info("declaring queue...")
	queue, err := ch.QueueDeclare("calendar_queue", true, false, false, false, nil)
	if err != nil {
		return err
	}
	log.Info("done.")
	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}
	log.Info("done.")
	q.conn = conn
	q.queue = queue
	q.ch = ch
	return nil
}

func (q *QueueClient) Close() error {
	log.Info("closing channel...")
	err := q.ch.Close()
	if err != nil {
		return err
	}
	log.Info("done.")
	return q.conn.Close()
}

func (q *QueueClient) SendMessage(message MessageDto) error {
	if q.isClosed() {
		return errors.New("rabbit mq is not available now.")
	}

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

func (q *QueueClient) ReceiveMessages() (<-chan WorkItem, error) {
	messageQueue, err := q.ch.Consume(q.queue.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	outputChannel := make(chan WorkItem)

	go func() {
		for {
			select {
			case d := <-messageQueue:
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
			case <-q.connectionClosedCh:
				log.Info("Receiver is going to be stopped")
				<-q.connectionEstablishedCh
				log.Info("Receiver is going to be resumed")
				messageQueue, _ = q.ch.Consume(q.queue.Name, "", false, false, false, false, nil)
			}
		}
	}()

	return outputChannel, nil
}

func (q *QueueClient) isClosed() bool {
	return atomic.LoadInt32(&q.closed) == 1
}

func (q *QueueClient) setIsClosed(value bool) {
	if value {
		atomic.StoreInt32(&q.closed, 1)
	} else {
		atomic.StoreInt32(&q.closed, 0)
	}
}
