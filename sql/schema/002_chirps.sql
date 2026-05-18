-- +goose Up
CREATE TABLE IF NOT EXISTS chirps(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    body TEXT NOT NULL,
    user_id UUID NOT NULL,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE chirps;