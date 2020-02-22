package errors

import (
	"fmt"
	"go-calenar/calendar/entities"
)

type ErrNotFound struct {
	message string
}

func (e ErrNotFound) Error() string {
	return e.message
}

func NewErrNotFound(id entities.Id) ErrNotFound {
	return ErrNotFound{message: fmt.Sprintf("unable to find event with the id %v", id)}
}
