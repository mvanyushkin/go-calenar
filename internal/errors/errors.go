package errors

import (
	"fmt"
	"github.com/mvanyushkin/go-calendar/internal/entities"
	"strings"
)

type ErrArgsInvalid struct {
	wrongArgs map[string]string
}

func NewErrArgsInvalid(wrongArgs map[string]string) ErrArgsInvalid {
	return ErrArgsInvalid{
		wrongArgs: wrongArgs,
	}
}

func (e ErrArgsInvalid) Error() string {
	sb := strings.Builder{}
	for s, s2 := range e.wrongArgs {
		sb.WriteString(fmt.Sprintf("The arg %v is incorrect: %v", s, s2))
	}
	return sb.String()
}

type ErrNotFound struct {
	message string
}

func (e ErrNotFound) Error() string {
	return e.message
}

func NewErrNotFound(id entities.Id) ErrNotFound {
	return ErrNotFound{message: fmt.Sprintf("unable to find event with the id %v", id)}
}

type ErrOutOfSchedule struct {
	message string
}

func (e ErrOutOfSchedule) Error() string {
	return e.message
}

type ErrTimeBusy struct {
	Message string
}

func (e ErrTimeBusy) Error() string {
	return e.Message
}
