package errors

type ErrTimeBusy struct {
	Message string
}

func (e ErrTimeBusy) Error() string {
	return e.Message
}
