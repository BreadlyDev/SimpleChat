CREATE TABLE IF NOT EXISTS profile(
    id SERIAL PRIMARY KEY,
    photo VARCHAR(1000),
    user_id INT,

    -- TODO: add new columns and indexes

    CONSTRAINT fk_user FOREIGN KEY user_id REFERENCES users(id)
);