package usecases

import (
	"github.com/stretchr/testify/assert"
	"go-calenar/calendar/entities"
	store "go-calenar/calendar/store"
	"testing"
	"time"
)

func TestWhenStoreHasElements(t *testing.T) {
	s := store.NewInMemoryEventStoreFromSlice([]entities.Event{{
		Id:          0,
		Title:       "",
		Description: "",
		Time:        time.Time{},
	}})

	useCase := FetchEventsUseCase{UseCase{store: s}}
	result, err := useCase.Do()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
}

func TestWhenStoreHasNoElements(t *testing.T) {
	s := store.NewInMemoryEventStore()
	useCase := FetchEventsUseCase{UseCase{store: s}}
	result, err := useCase.Do()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}
