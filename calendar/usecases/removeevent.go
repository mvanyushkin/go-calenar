package usecases

import (
	"github.com/thoas/go-funk"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/errors"
)

type RemoveEventUseCase struct {
	UseCase
	Event *entities.Event
}

func (u *RemoveEventUseCase) Do() error {
	events, _ := u.store.List()
	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Id == u.Event.Id
	}).([]entities.Event)

	if len(values) == 0 {
		return errors.NewErrNotFound(u.Event.Id)
	}

	return u.store.Remove(u.Event)
}
