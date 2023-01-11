FROM postgres

ENV POSTGRES_USER=finances@ruben
ENV POSTGRES_PASSWORD=rub3nF!n4nc3s

COPY ./schema/0001-create-table-expenses.up.sql /docker-entrypoint-initdb.d/0001-create-table-expenses.up.sql
COPY ./data/1001-expense-categories.sql /docker-entrypoint-initdb.d/1001-expense-categories.sql
COPY ./data/1002-expense-subcategories.sql /docker-entrypoint-initdb.d/1002-expense-subcategories.sql