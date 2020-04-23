package reminder

import (
	"context"
	"errors"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/messages"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type reminder struct {
	db                    *sqlx.DB
	dbConnectionString    string
	queueConnectionString string
	queueClient           *messages.QueueClient
	context               context.Context
}

func New(ctx context.Context, dbConnectionString string, queueConnectionString string) (*reminder, error) {
	r := &reminder{
		dbConnectionString:    dbConnectionString,
		queueConnectionString: queueConnectionString,
		context:               ctx,
	}
	err := r.openSession()
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *reminder) Do() error {
	timer := time.NewTimer(time.Nanosecond)
	for {
		select {
		case <-timer.C:
			err := r.internalDo()
			if err != nil {
				log.Errorf("Current reminding iteration failed, occurred an exception: %v", err.Error())
			}
			log.Info("Sleeping.... Z-z-z")
			time.Sleep(time.Minute)
			log.Info("WAKING UP!")
		case <-r.context.Done():
			return nil
		}
		timer = time.NewTimer(time.Minute)
	}
}

func (r *reminder) Close() {
	r.closeSession()
}

func (r *reminder) internalDo() error {
	rows, err := r.selectRows()
	if err != nil {
		return err
	}

	for _, row := range rows {
		tx, err := r.db.Begin()
		if err != nil {
			return err
		}

		err = r.setReminded(row)
		if err != nil {
			rbackErr := tx.Rollback()
			if rbackErr != nil {
				return errors.New(strings.Join([]string{rbackErr.Error(), err.Error()}, ","))
			}
			return err
		}

		err = r.sendToQueue(row)
		if err != nil {
			rbackErr := tx.Rollback()
			if rbackErr != nil {
				return errors.New(strings.Join([]string{rbackErr.Error(), err.Error()}, ","))
			}
			return err
		}

		err = tx.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *reminder) openSession() error {
	log.Info("Opening database connection...")
	db, err := sqlx.Open("pgx", r.dbConnectionString)
	if err != nil {
		return err
	}
	r.db = db
	log.Info("Done.")

	log.Info("Opening queue client connection...")
	c, err := messages.NewQueueClient(r.queueConnectionString)
	if err != nil {
		return err
	}
	log.Info("Done.")

	r.queueClient = c
	return nil
}

func (r *reminder) closeSession() {
	log.Info("Closing database connection...")
	if r.db != nil {
		_ = r.db.Close()
		r.db = nil
	}
	log.Info("Done.")

	log.Info("Closing database connection...")
	r.queueClient.Close()
	log.Info("Done.")
}

func (r *reminder) setReminded(e entities.Event) error {
	log.Infof("Setting the reminded field as true for the event id %v ...", e.Id)
	args := map[string]interface{}{"event_id": e.Id}
	_, err := r.db.NamedExec("UPDATE events SET reminded = true WHERE id = :event_id", args)
	if err != nil {
		return err
	}
	log.Info("Done.")
	return err
}

func (r *reminder) sendToQueue(e entities.Event) error {
	log.Infof("Sending an event to the queue %v ...", e.Id)

	err := r.queueClient.SendMessage(messages.MessageDto{
		Title:       string(e.Title),
		Description: string(e.Description),
	})
	if err != nil {
		return err
	}

	log.Info("Done.")
	return nil
}

func (r *reminder) selectRows() ([]entities.Event, error) {
	tomorrow := time.Now().AddDate(0, 0, 1)

	log.Info("Querying rows...")
	q, err := r.db.Queryx("SELECT id, title, description from events WHERE date < $1 LIMIT 1000", tomorrow)
	if err != nil {
		return nil, err
	}

	events := make([]entities.Event, 0)
	for q.Next() {
		event := entities.Event{}
		err := q.StructScan(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	log.Info("Done.")
	err = q.Close()
	if err != nil {
		return nil, err
	}

	return events, nil
}
