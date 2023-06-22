CREATE TABLE if not exists movies(
    id bigserial PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    year INTEGER NOT NULL,
    runtime INTEGER NOT NULL,
    genres text[] NOT NULL,
    version INTEGER NOT NULL DEFAULT 1
);
