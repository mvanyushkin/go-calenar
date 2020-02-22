package calendar

import (
	"github.com/mvanyushkin/go-calendar/calendar/entities"
	"github.com/mvanyushkin/go-calendar/calendar/store"
	"github.com/mvanyushkin/go-calendar/calendar/usecases"
	"time"
)

type calendar struct {
	store *store.EventStore
}

func NewCalendar(store store.EventStore) calendar {
	return calendar{store: &store}
}

func NewCal(store store.EventStore) calendar {
	return calendar{store: &store}
}

func (c *calendar) CreateEvent(title entities.Title, description entities.Description, desiredTime time.Time) (*entities.Event, error) {
	useCase := usecases.CreateEventUseCase{
		UseCase:     usecases.NewBaseUseCase(c.store),
		Title:       title,
		Description: description,
		Time:        desiredTime,
	}

	return useCase.Do()
}

func (c *calendar) Remove(event *entities.Event) error {
	useCase := usecases.RemoveEventUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
		Event:   event,
	}

	return useCase.Do()
}

func (c *calendar) List() ([]entities.Event, error) {
	useCase := usecases.FetchEventsUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
	}

	return useCase.Do()
}

func (c *calendar) Get(id entities.Id) (*entities.Event, error) {
	useCase := usecases.GetEventUseCase{
		UseCase: usecases.NewBaseUseCase(c.store),
		Id:      id,
	}

	return useCase.Do()
}
