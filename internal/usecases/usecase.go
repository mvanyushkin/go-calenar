package usecases

import "github.com/mvanyushkin/go-calendar/internal/store"

type UseCase struct {
	store store.EventStore
}

func NewBaseUseCase(store *store.EventStore) UseCase {
	return UseCase{store: *store}
}
