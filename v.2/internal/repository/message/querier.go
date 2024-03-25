package message

import (
	"chat-lesson/internal/model"

	"context"
)

type Querier interface {
	MessageCreate(ctx context.Context, arg model.MessageNew) ([]byte, error)
}

var _ Querier = (*Queries)(nil)
