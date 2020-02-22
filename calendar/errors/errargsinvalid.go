package errors

import (
	"fmt"
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
