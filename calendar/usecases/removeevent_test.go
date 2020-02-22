package usecases

import (
	"github.com/stretchr/testify/assert"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/store"
	"testing"
	"time"
)

func TestWhenRemovingEventExists(t *testing.T) {
	event := entities.Event{
		Id:          111,
		Title:       "",
		Description: "",
		Time:        time.Time{},
	}
	s := store.NewInMemoryEventStoreFromSlice([]entities.Event{
		event,
	})

	u := RemoveEventUseCase{
		UseCase: UseCase{store: s},
		Event:   &event,
	}

	err := u.Do()
	assert.Nil(t, err)
}

func TestWhenRemovingEventDoesntExist(t *testing.T) {
	s := store.NewInMemoryEventStore()
	event := entities.Event{
		Id:          111,
		Title:       "",
		Description: "",
		Time:        time.Time{},
	}
	u := RemoveEventUseCase{
		UseCase: UseCase{store: s},
		Event:   &event,
	}

	err := u.Do()
	assert.NotNil(t, err)
}
