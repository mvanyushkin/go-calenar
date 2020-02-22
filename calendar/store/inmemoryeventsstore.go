package store

import (
	"github.com/thoas/go-funk"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/errors"
	"sync"
)

var globalCounter int
var globalCounterMutex = sync.Mutex{}

func getNextId() entities.Id {
	globalCounterMutex.Lock()
	defer globalCounterMutex.Unlock()
	globalCounter++
	return entities.Id(globalCounter)
}

type inMemoryEventStore struct {
	events []entities.Event
}

func (s *inMemoryEventStore) Update(event *entities.Event) error {
	// Do nothing, 'cause it returns object by a pointer so we work with single object every time
	return nil
}

func NewInMemoryEventStore() *inMemoryEventStore {
	return &inMemoryEventStore{events: make([]entities.Event, 0)}
}

func NewInMemoryEventStoreFromSlice(events []entities.Event) *inMemoryEventStore {
	return &inMemoryEventStore{events: events}
}

func (s *inMemoryEventStore) Add(event *entities.Event) error {
	event.Id = getNextId()
	s.events = append(s.events, *event)
	return nil
}

func (s *inMemoryEventStore) Remove(event *entities.Event) error {
	s.events = funk.Filter(s.events, func(x entities.Event) bool {
		return x.Id != event.Id
	}).([]entities.Event)
	return nil
}

func (s *inMemoryEventStore) List() ([]entities.Event, error) {
	return s.events, nil
}

func (s *inMemoryEventStore) Get(id entities.Id) (*entities.Event, error) {
	if id < 1 {
		return nil, errors.NewErrArgsInvalid(map[string]string{"id": "the specified id must be greater than zero"})
	}

	foundItems := funk.Filter(s.events, func(x entities.Event) bool {
		return x.Id == id
	}).([]entities.Event)

	if len(foundItems) == 0 {
		return nil, errors.NewErrNotFound(id)
	}

	return &foundItems[0], nil
}
