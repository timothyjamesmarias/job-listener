-- +goose Up
-- +goose StatementBegin
    CREATE TABLE jobs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
        app_id INTEGER,
        FOREIGN KEY (app_id) REFERENCES apps(id)
    );
    CREATE INDEX jobs_id_index ON jobs (id);
    CREATE INDEX jobs_app_id_index ON jobs (app_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE jobs;
    DROP INDEX jobs_id_index;
-- +goose StatementEnd
