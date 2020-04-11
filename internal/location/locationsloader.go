package location

import "github.com/mvanyushkin/go-calendar/internal/entities"

type LocationsLoader interface {
	Do() ([]entities.Location, error)
}
