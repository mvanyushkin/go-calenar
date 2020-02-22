package store

import (
	"go-calenar/calendar/entities"
)

type EventStore interface {
	Add(event *entities.Event) error
	Remove(event *entities.Event) error
	Update(event *entities.Event) error
	List() ([]entities.Event, error)
	Get(id entities.Id) (*entities.Event, error)
}