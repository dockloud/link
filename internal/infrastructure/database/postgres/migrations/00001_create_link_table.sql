-- +goose Up
CREATE TABLE links (
    token VARCHAR(120) PRIMARY KEY,
    document_id INT NOT NULL,
    creator_id UUID NOT NULL,

    read_allowed BOOLEAN NOT NULL DEFAULT TRUE,
    write_allowed BOOLEAN NOT NULL DEFAULT FALSE,
    delete_allowed BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ,
    revoked_at TIMESTAMPTZ,

    CHECK(read_allowed OR write_allowed OR delete_allowed),
    CHECK(expires_at IS NULL OR expires_at > created_at)
);

-- +goose Down
DROP TABLE IF EXISTS links;
