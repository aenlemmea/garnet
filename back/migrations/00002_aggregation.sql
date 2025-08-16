-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Aggregation (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    blurb TEXT,
    link TEXT UNIQUE NOT NULL,
    origin_name TEXT NOT NULL,
    tags TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Aggregation;
-- +goose StatementEnd