package usecases

import (
	"github.com/thoas/go-funk"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/errors"
	"time"
)

type CreateEventUseCase struct {
	UseCase
	Title       entities.Title
	Description entities.Description
	Time        time.Time
}

func (u *CreateEventUseCase) Do() (*entities.Event, error) {
	validationErrors := validateArgs(u)
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

	err = u.store.Add(newEvent)
	if err != nil {
		return nil, err
	}

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

func validateArgs(useCase *CreateEventUseCase) map[string]string {
	errorsBag := make(map[string]string)
	if useCase.Title == "" {
		errorsBag["title"] = "cannot be null or empty"
	}

	if useCase.Description == "" {
		errorsBag["description"] = "cannot be null or empty"
	}

	return errorsBag
}
