package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	store "github.com/mvanyushkin/go-calendar/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWhenStoreHasElements(t *testing.T) {
	s := store.NewInMemoryEventStore().LoadFromSlice([]entities.Event{{
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
