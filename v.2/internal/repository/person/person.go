package person

import (
	"chat-lesson/internal/model"
	"context"
	"log"
)

const createPersonSQL = `-- name: PersonCreate :one
insert into persons(username) values($1);`

func (q *Queries) PersonCreate(ctx context.Context, arg model.PersonNew) ([]byte, error) {
	log.Println("PersonCreate")
	row := q.db.QueryRow(ctx, createPersonSQL,
		arg.UserName,
	)
	result := []byte(``)
	err := row.Scan(&result)
	return result, err
}
