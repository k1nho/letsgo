CREATE TABLE IF NOT EXISTS tokens(
    hash BYTEA PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    expiry TIMESTAMP(0) WITH time zone NOT NULL,
    scope TEXT NOT NULL
);
