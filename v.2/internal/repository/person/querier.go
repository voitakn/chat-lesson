package person

import (
	"chat-lesson/internal/model"

	"context"
)

type Querier interface {
	PersonCreate(ctx context.Context, arg model.PersonNew) ([]byte, error)
}

var _ Querier = (*Queries)(nil)
