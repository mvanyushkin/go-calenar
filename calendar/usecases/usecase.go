package usecases

import "github.com/mvanyushkin/go-calendar/calendar/store"

type UseCase struct {
	store store.EventStore
}

func NewBaseUseCase(store *store.EventStore) UseCase {
	return UseCase{store: *store}
}
