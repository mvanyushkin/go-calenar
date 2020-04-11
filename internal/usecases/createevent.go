package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/errors"
	"github.com/thoas/go-funk"
	"time"
)

type CreateEventUseCase struct {
	UseCase
	Title       entities.Title
	Description entities.Description
	Time        time.Time
}

func (u *CreateEventUseCase) Do() (*entities.Event, error) {
	validationErrors := validateCreateArgs(u)
	if len(validationErrors) > 0 {
		return nil, errors.NewErrArgsInvalid(validationErrors)
	}

	events, err := u.store.List()
	if err != nil {
		return nil, err
	}

	event, errTimeBusy := u.checkIfTimeIsBusy(events)
	if errTimeBusy != nil {
		return event, errTimeBusy
	}

	newEvent := &entities.Event{
		Title:       u.Title,
		Description: u.Description,
		Time:        u.Time,
	}

	id, err := u.store.Add(newEvent)
	if err != nil {
		return nil, err
	}

	newEvent.Id = id
	return newEvent, err
}

func (u *CreateEventUseCase) checkIfTimeIsBusy(events []entities.Event) (*entities.Event, error) {
	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Time == u.Time
	}).([]entities.Event)

	if len(values) > 0 {
		return nil, errors.ErrTimeBusy{Message: "The specified time busy."}
	}
	return nil, nil
}

func validateCreateArgs(useCase *CreateEventUseCase) map[string]string {
	errorsBag := make(map[string]string)
	if useCase.Title == "" {
		errorsBag["title"] = "cannot be null or empty"
	}

	if useCase.Description == "" {
		errorsBag["description"] = "cannot be null or empty"
	}

	return errorsBag
}
