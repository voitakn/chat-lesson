create table if not exists persons (
    id uuid not null primary key default new_uuid(),
    username varchar(40)
);

create table if not exists messages (
    id uuid not null primary key default new_uuid(),
    person_id uuid references persons(id),
    created_at timestamptz default now(),
    body text
);