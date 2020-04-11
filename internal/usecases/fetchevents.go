package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
)

type FetchEventsUseCase struct {
	UseCase
}

func (u *FetchEventsUseCase) Do() ([]entities.Event, error) {
	return u.store.List()
}
