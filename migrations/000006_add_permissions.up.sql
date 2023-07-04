CREATE TABLE IF NOT EXISTS permissions(
    id bigserial PRIMARY KEY,
    code TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_permissions(
    user_id bigserial NOT NULL REFERENCES users ON DELETE CASCADE,
    permission_id bigserial NOT NULL REFERENCES permissions ON DELETE CASCADE,
    PRIMARY KEY(user_id, permission_id)
);

INSERT INTO permissions(code)
VALUES('movies:read'), ('movies:write');
