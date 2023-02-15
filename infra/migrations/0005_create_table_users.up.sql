CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(20) NOT NULL UNIQUE,
    passhash VARCHAR(60)
);

INSERT INTO users (username, passhash) VALUES ('rubengomes8', '$2a$10$p3MNALB71zkMTm.m0eCBHOO.b9I8dwoeJ6/698peFyh.DIoYgjhwS');