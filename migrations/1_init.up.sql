CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255),
    pass_hash BYTEA
);

CREATE INDEX idx_unique_email ON users(email);