package usecases

import (
	"github.com/stretchr/testify/assert"
	"go-calenar/calendar/entities"
	"go-calenar/calendar/errors"
	"go-calenar/calendar/store"
	"testing"
	"time"
)

func TestReappointEventWhenTimeIsFree(t *testing.T) {
	s := store.NewInMemoryEventStore()
	evt := entities.Event{
		Id:          1,
		Title:       "",
		Description: "",
		Time:        time.Now(),
	}
	useCase := UserReappointEvent{
		UseCase:       UseCase{store: s},
		DesiredTime:   time.Now(),
		ExistingEvent: evt,
	}

	err := useCase.Do()
	assert.Nil(t, err)
}

func TestReappointEventWhenTimeIsBusy(t *testing.T) {
	var busyTime = time.Now()
	var s = store.NewInMemoryEventStoreFromSlice([]entities.Event{{
		Id:          1,
		Title:       "",
		Description: "",
		Time:        busyTime,
	}})

	evt := entities.Event{
		Id:          1,
		Title:       "",
		Description: "",
		Time:        time.Now(),
	}

	useCase := UserReappointEvent{
		UseCase:       UseCase{store: s},
		DesiredTime:   busyTime,
		ExistingEvent: evt,
	}

	err := useCase.Do()
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrTimeBusy{})
}
