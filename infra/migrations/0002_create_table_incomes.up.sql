CREATE TABLE income_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20)
);

CREATE TABLE incomes (
    id SERIAL PRIMARY KEY,
    value FLOAT NOT NULL,
    date DATE NOT NULL,
    description VARCHAR(50),

    category_id INTEGER NOT NULL,
    card_id INTEGER NOT NULL,

    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES income_categories(id),
    CONSTRAINT fk_card FOREIGN KEY(card_id) REFERENCES cards(id)

);

INSERT INTO income_categories (name) values ('Salary');
INSERT INTO income_categories (name) values ('DeGiro');
INSERT INTO income_categories (name) values ('BetClic');