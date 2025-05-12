-- Schema
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Seed data
INSERT INTO users (email) VALUES
('test1@example.com'),
('test2@example.com');
