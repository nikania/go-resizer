CREATE TABLE IF NOT EXISTS  users (
id serial PRIMARY KEY,
login VARCHAR(30) UNIQUE NOT NULL,
password_hash bytea NOT NULL,
email VARCHAR(300) UNIQUE NOT NULL
);