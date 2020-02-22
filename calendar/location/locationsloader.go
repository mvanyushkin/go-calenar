package location

import "github.com/mvanyushkin/go-calendar/calendar/entities"

type LocationsLoader interface {
	Do() ([]entities.Location, error)
}
