-- +goose Up
-- +goose StatementBegin
    CREATE TABLE jobs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );
    CREATE INDEX jobs_id_index ON apps (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE jobs;
    DROP INDEX jobs_id_index;
-- +goose StatementEnd
