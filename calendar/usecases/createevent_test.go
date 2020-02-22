package usecases

import (
	"github.com/stretchr/testify/assert"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/store"
	"testing"
	"time"
)

func TestCreateNewWhenTimeIsFree(t *testing.T) {
	s := store.NewInMemoryEventStore()
	useCase := CreateEventUseCase{
		UseCase:     UseCase{store: s},
		Title:       "11111",
		Description: "11111",
		Time:        time.Now(),
	}

	evt, err := useCase.Do()
	assert.Nil(t, err)
	assert.NotNil(t, evt)
	assert.NotEqual(t, evt.Id, 0)
}

func TestCreateNewWhenTimeIsBusy(t *testing.T) {
	var busyTime = time.Now()
	var s = store.NewInMemoryEventStoreFromSlice([]entities.Event{{
		Id:          1,
		Title:       "",
		Description: "",
		Time:        busyTime,
	}})
	useCase := CreateEventUseCase{
		UseCase:     UseCase{store: s},
		Title:       "",
		Description: "",
		Time:        busyTime,
	}

	evt, err := useCase.Do()
	assert.Nil(t, evt)
	assert.NotNil(t, err)
}
