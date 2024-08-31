-- +gooseUp

-- +gooseStatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);
-- +gooseStatementEnd

-- +gooseStatementBegin
CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    content TEXT
)
-- +gooseStatementEnd

-- +gooseDown
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS notes;