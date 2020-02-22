package usecases

import (
	"go-calenar/calendar/entities"
)

type FetchEventsUseCase struct {
	UseCase
}

func (u *FetchEventsUseCase) Do() ([]entities.Event, error) {
	return u.store.List()
}
