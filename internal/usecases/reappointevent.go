package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/errors"
	"github.com/thoas/go-funk"
	"time"
)

type UserReappointEvent struct {
	UseCase
	ExistingEvent entities.Event
	DesiredTime   time.Time
}

func (u *UserReappointEvent) Do() error {
	events, err := u.store.List()
	if err != nil {
		return err
	}

	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Time == u.DesiredTime
	}).([]entities.Event)

	if len(values) > 0 {
		return errors.ErrTimeBusy{Message: "The specified time busy."}
	}

	return u.store.Update(&u.ExistingEvent)
}
