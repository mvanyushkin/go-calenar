package internal

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/store"
	"github.com/mvanyushkin/go-calendar/internal/usecases"
	"time"
)

type Calendar struct {
	store *store.EventStore
}

func NewCalendar(store store.EventStore) *Calendar {
	return &Calendar{store: &store}
}

func (c *Calendar) CreateEvent(title entities.Title, description entities.Description, desiredTime time.Time) (*entities.Event, error) {
	useCase := usecases.CreateEventUseCase{
		UseCase:     usecases.NewBaseUseCase(c.store),
		Title:       title,
		Description: description,
		Time:        desiredTime,
	}

	return useCase.Do()
}

func (c *Calendar) Remove(event *entities.Event) error {
	useCase := usecases.RemoveEventUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
		Event:   event,
	}

	return useCase.Do()
}

func (c *Calendar) List() ([]entities.Event, error) {
	useCase := usecases.FetchEventsUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
	}

	return useCase.Do()
}

func (c *Calendar) Get(id entities.Id) (*entities.Event, error) {
	useCase := usecases.GetEventUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
		Id:      id,
	}

	return useCase.Do()
}

func (c *Calendar) Update(evt entities.Event) error {
	useCase := usecases.UpdateEventUseCase{
		UseCase:     usecases.NewBaseUseCase(c.store),
		Id:          evt.Id,
		Title:       evt.Title,
		Description: evt.Description,
		Time:        evt.Time,
	}

	return useCase.Do()
}
