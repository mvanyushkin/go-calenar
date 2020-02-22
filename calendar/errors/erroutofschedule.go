package errors

type ErrOutOfSchedule struct {
	message string
}

func (e ErrOutOfSchedule) Error() string {
	return e.message
}
