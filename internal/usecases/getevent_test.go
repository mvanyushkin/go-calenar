package usecases

import (
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"github.com/mvanyushkin/go-calendar/internal/errors"
	"github.com/mvanyushkin/go-calendar/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWhenIdIsCorrectAndEventExists(t *testing.T) {
	var testId entities.Id = 999
	s := store.NewInMemoryEventStore().LoadFromSlice([]entities.Event{
		{Id: testId, Title: "", Description: "", Time: time.Now()},
	})

	event, err := s.Get(testId)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, event.Id, testId)
}

func TestWhenIdIsIncorrect(t *testing.T) {
	s := store.NewInMemoryEventStore()

	event, err := s.Get(-1)
	assert.Nil(t, event)
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.NewErrArgsInvalid(map[string]string{}))
}

func TestWhenEventDoesntExist(t *testing.T) {
	s := store.NewInMemoryEventStore()

	event, err := s.Get(999)
	assert.Nil(t, event)
	assert.NotNil(t, err)
	assert.IsType(t, err, errors.ErrNotFound{})
}
