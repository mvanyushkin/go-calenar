package location

import "go-calenar/calendar/entities"

type LocationsLoader interface {
	Do() ([]entities.Location, error)
}
