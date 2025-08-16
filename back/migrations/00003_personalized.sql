-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Personalized (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    aggregation_id BIGINT NOT NULL REFERENCES Aggregation(id) ON DELETE CASCADE,
    score FLOAT,
    sent BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, aggregation_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Personalized;
-- +goose StatementEnd