package sender

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/mvanyushkin/go-calendar/internal/messages"
)

type sender struct {
	queueConnectionString string
}

func CreateSender(queueConnectionString string) *sender {
	return &sender{
		queueConnectionString: queueConnectionString,
	}
}

func (s sender) ListenMessages() error {

	queueClient, err := messages.NewQueueClient(s.queueConnectionString)
	if err != nil {
		return err
	}

	incomingCh, err := queueClient.ReceiveMessages()
	if err != nil {
		return nil
	}

	for workItem := range incomingCh {
		fmt.Printf("Reminding: %v %v \n", workItem.Message.Title, workItem.Message.Description)
		workItem.DoneCallback()
	}

	return nil
}
