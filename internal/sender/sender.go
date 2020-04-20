package sender

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/mvanyushkin/go-calendar/internal/messages"
)

type sender struct {
	queueConnectionString string
	ctx                   context.Context
}

func CreateSender(queueConnectionString string, ctx context.Context) *sender {
	return &sender{
		queueConnectionString: queueConnectionString,
		ctx:                   ctx,
	}
}

func (s *sender) ListenMessages() error {

	queueClient, err := messages.NewQueueClient(s.queueConnectionString)
	if err != nil {
		return err
	}

	incomingCh, err := queueClient.ReceiveMessages()
	if err != nil {
		return nil
	}

	for {
		select {
		case workItem := <-incomingCh:
			fmt.Printf("Reminding: %v %v \n", workItem.Message.Title, workItem.Message.Description)
			workItem.DoneCallback()
		case <-s.ctx.Done():
			return nil
		}
	}
}
