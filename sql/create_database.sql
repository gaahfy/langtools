CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) UNIQUE NOT NULL CHECK (username ~ '^[A-Za-z0-9]{4,20}$'),
    password_hash CHAR(64) NOT NULL,
    creation_time TIMESTAMP DEFAULT NOW(),
    ip_hash CHAR(64) NOT NULL,
    login_attempts INT NOT NULL DEFAULT 0,
    last_login_attempt TIMESTAMP
);