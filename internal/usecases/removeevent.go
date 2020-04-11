package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/errors"
	"github.com/thoas/go-funk"
)

type RemoveEventUseCase struct {
	UseCase
	Event *entities.Event
}

func (u *RemoveEventUseCase) Do() error {
	events, err := u.store.List()
	if err != nil {
		return err
	}

	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Id == u.Event.Id
	}).([]entities.Event)

	if len(values) == 0 {
		return errors.NewErrNotFound(u.Event.Id)
	}

	return u.store.Remove(u.Event)
}
