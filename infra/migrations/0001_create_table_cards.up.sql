CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE
);

INSERT INTO cards (name) values ('CGD');
INSERT INTO cards (name) values ('Food allowance');