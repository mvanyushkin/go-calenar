package usecases

import (
	"github.com/mvanyushkin/go-calendar/calendar/entities"
	"github.com/mvanyushkin/go-calendar/calendar/errors"
)

type GetEventUseCase struct {
	UseCase
	Id entities.Id
}

func (u *GetEventUseCase) Do() (*entities.Event, error) {
	if u.Id <= 0 {
		return nil, errors.NewErrArgsInvalid(map[string]string{
			"Invalid Id value": "Must be a positive value",
		})
	}

	event, err := u.store.Get(u.Id)
	if event == nil {
		return nil, errors.NewErrNotFound(u.Id)
	}

	return event, err
}
