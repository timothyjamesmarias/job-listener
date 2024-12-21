-- +goose Up
-- +goose StatementBegin
    CREATE TABLE apps (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        id_hash varchar(255) NOT NULL,
        name varchar(255) NOT NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL,
    );

    CREATE INDEX id_hash_index ON apps (id_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE apps;
    DROP INDEX id_hash_index;
-- +goose StatementEnd
