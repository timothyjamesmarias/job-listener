-- +goose Up
-- +goose StatementBegin
    CREATE TABLE apps (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name varchar(255) NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    );
    CREATE INDEX apps_id_index ON apps (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE apps;
    DROP INDEX apps_id_index;
-- +goose StatementEnd
