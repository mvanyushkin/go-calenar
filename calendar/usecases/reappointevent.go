package usecases

import (
	"github.com/thoas/go-funk"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/errors"
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
