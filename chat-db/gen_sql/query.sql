-- name: CreateUser :one
insert into persons(username) VALUES ($1)
returning id, username;

-- name: CheckUser :one
select id, username from persons
where id = $1;

-- name: Users :many
select id, username from persons
limit $1 offset $2;

-- name: CreateMessage :one
insert into messages(person_id, body)
VALUES ($1, $2)
returning id, person_id, created_at, body;

-- name: Messages :many
select id, person_id, created_at, body from messages
limit $1 offset $2;

-- name: MessagesByPerson :many
select id, person_id, created_at, body
from messages
where person_id = $1
limit $2 offset $3;
