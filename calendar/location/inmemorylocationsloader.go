package location

import (
	"github.com/mvanyushkin/go-calendar/calendar/entities"
)

type InMemoryLocationsLoader struct{}

var hardcodedLocations = []entities.Location{
	{RoomNumber: 1},
	{RoomNumber: 2},
	{RoomNumber: 3},
	{RoomNumber: 4},
	{RoomNumber: 5},
}

func (InMemoryLocationsLoader) Do() ([]entities.Location, error) {
	return hardcodedLocations, nil
}
