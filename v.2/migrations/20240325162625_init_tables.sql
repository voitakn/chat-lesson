-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION new_ulid() RETURNS text
AS $$
SELECT lpad(to_hex(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint), 12, '0')
           || encode(gen_random_bytes(10), 'hex');
$$ LANGUAGE SQL;

CREATE OR REPLACE FUNCTION new_uuid() RETURNS uuid
AS $$
SELECT new_ulid()::uuid;
$$ LANGUAGE SQL;

create table if not exists persons
(
    id uuid not null primary key default new_uuid(),
    username varchar(40)
);

create table if not exists messages
(
    id uuid not null primary key default new_uuid(),
    person_id uuid references persons(id),
    created_at timestamptz default now(),
    body text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists messages;
drop table if exists persons;
-- +goose StatementEnd
