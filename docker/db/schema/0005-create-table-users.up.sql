CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    salt VARCHAR(16),
    passhash VARSHAR(40)
);