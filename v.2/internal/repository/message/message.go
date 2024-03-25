package message

import (
	"chat-lesson/internal/model"
	"context"
)

func (q *Queries) MessageCreate(ctx context.Context, arg model.MessageNew) ([]byte, error) {
	return nil, nil
}
