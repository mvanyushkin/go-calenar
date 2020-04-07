package usecases

import (
	"github.com/mvanyushkin/go-calendar/calendar/entities"
	"github.com/mvanyushkin/go-calendar/calendar/errors"
	"github.com/thoas/go-funk"
	"time"
)

type UpdateEventUseCase struct {
	UseCase
	Id          entities.Id
	Title       entities.Title
	Description entities.Description
	Time        time.Time
}

func (u *UpdateEventUseCase) Do() error {
	validationErrors := validateUpdateArgs(u)
	if len(validationErrors) > 0 {
		return errors.NewErrArgsInvalid(validationErrors)
	}

	events, err := u.store.List()
	if err != nil {
		return err
	}

	errTimeBusy := u.checkIfTimeIsBusy(events)
	if errTimeBusy != nil {
		return errTimeBusy
	}

	event, err := u.store.Get(u.Id)
	if err != nil {
		return err
	}

	event.Time = u.Time
	event.Title = u.Title
	event.Description = u.Description

	return u.store.Update(event)
}

func (u *UpdateEventUseCase) checkIfTimeIsBusy(events []entities.Event) error {
	var values = funk.Filter(events, func(x entities.Event) bool {
		return x.Time == u.Time && x.Id != u.Id
	}).([]entities.Event)

	if len(values) > 0 {
		return errors.ErrTimeBusy{Message: "The specified time busy."}
	}
	return nil
}

func validateUpdateArgs(useCase *UpdateEventUseCase) map[string]string {
	errorsBag := make(map[string]string)
	if useCase.Title == "" {
		errorsBag["title"] = "cannot be null or empty"
	}

	if useCase.Description == "" {
		errorsBag["description"] = "cannot be null or empty"
	}

	return errorsBag
}
