package usecases

import "go-calenar/calendar/store"

type UseCase struct {
	store store.EventStore
}

func NewBaseUseCase(store *store.EventStore) UseCase {
	return UseCase{store: *store}
}
