CREATE TABLE expense_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE expense_subcategories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    category_id INTEGER NOT NULL,

    CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES expense_categories(id)
);

CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    value FLOAT NOT NULL,
    date DATE NOT NULL,
    description VARCHAR(50),

    subcategory_id INTEGER NOT NULL,
    card_id INTEGER NOT NULL,

    CONSTRAINT fk_subcategory FOREIGN KEY(subcategory_id) REFERENCES expense_subcategories(id),
    CONSTRAINT fk_card FOREIGN KEY(card_id) REFERENCES cards(id)
);

INSERT INTO expense_categories (name) values ('House');
INSERT INTO expense_categories (name) values ('Health');
INSERT INTO expense_categories (name) values ('Transportation');
INSERT INTO expense_categories (name) values ('Leisure');
INSERT INTO expense_categories (name) values ('Investments');
INSERT INTO expense_categories (name) values ('Personal');

/* HOUSE subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('Rent', 1);
INSERT INTO expense_subcategories (name, category_id) values ('Furniture', 1);
INSERT INTO expense_subcategories (name, category_id) values ('Supermarket', 1);

/* HEALTH subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('Blood tests', 2);
INSERT INTO expense_subcategories (name, category_id) values ('Medical exam', 2);
INSERT INTO expense_subcategories (name, category_id) values ('Medical consult', 2);

/* TRANSPORTATION subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('Fuel', 3);
INSERT INTO expense_subcategories (name, category_id) values ('Highway', 3);
INSERT INTO expense_subcategories (name, category_id) values ('Bus', 3);
INSERT INTO expense_subcategories (name, category_id) values ('Uber', 3);
INSERT INTO expense_subcategories (name, category_id) values ('Airplane', 3);

/* LEISURE subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('Restaurants', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Uber Eats', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Books', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Cinema', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Hotel', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Hostel', 4);
INSERT INTO expense_subcategories (name, category_id) values ('Others', 4);

/* INVESTMENTS subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('DeGiro', 5);
INSERT INTO expense_subcategories (name, category_id) values ('BetClic', 5);

/* PERSONAL subcategories */
INSERT INTO expense_subcategories (name, category_id) values ('Barber', 6);
INSERT INTO expense_subcategories (name, category_id) values ('Clothes', 6);
INSERT INTO expense_subcategories (name, category_id) values ('Mobile Phone', 6);
INSERT INTO expense_subcategories (name, category_id) values ('Spotify', 6);
INSERT INTO expense_subcategories (name, category_id) values ('Benfica', 6);
INSERT INTO expense_subcategories (name, category_id) values ('MB', 6);
