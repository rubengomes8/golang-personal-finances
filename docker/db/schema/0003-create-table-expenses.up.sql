CREATE TABLE expense_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20)
);

CREATE TABLE expense_subcategories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    category_id INTEGER,

    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES expense_categories(id)
);

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    value FLOAT NOT NULL,
    date DATE NOT NULL,
    description VARCHAR(50),

    category_id INTEGER NOT NULL,
    subcategory_id INTEGER NOT NULL,
    card_id INTEGER NOT NULL,

    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES expense_categories(id),
    CONSTRAINT fk_subcategory FOREIGN KEY(subcategory_id) REFERENCES expense_subcategories(id),
    CONSTRAINT fk_card FOREIGN KEY(card_id) REFERENCES cards(id)
);