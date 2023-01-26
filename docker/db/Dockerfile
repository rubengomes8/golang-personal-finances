FROM postgres

ENV POSTGRES_USER=finances@ruben
ENV POSTGRES_PASSWORD=rub3nF!n4nc3s
ENV POSTGRES_DB=finances

COPY ./schema/0001-create-table-cards.up.sql /docker-entrypoint-initdb.d/0001-create-table-cards.up.sql
COPY ./schema/0002-create-table-incomes.up.sql /docker-entrypoint-initdb.d/0002-create-table-incomes.up.sql
COPY ./schema/0003-create-table-expenses.up.sql /docker-entrypoint-initdb.d/0003-create-table-expenses.up.sql
COPY ./schema/0004-create-view-expenses.up.sql /docker-entrypoint-initdb.d/0004-create-view-expenses.up.sql
COPY ./schema/0005-create-table-users.up.sql /docker-entrypoint-initdb.d/0005-create-table-users.up.sql
COPY ./schema/0006-create-view-incomes.up.sql /docker-entrypoint-initdb.d/0006-create-view-incomes.up.sql

COPY ./data/1001-cards.sql /docker-entrypoint-initdb.d/1001-cards.sql
COPY ./data/1002-expense-categories.sql /docker-entrypoint-initdb.d/1002-expense-categories.sql
COPY ./data/1003-expense-subcategories.sql /docker-entrypoint-initdb.d/1003-expense-subcategories.sql
COPY ./data/1004-income-categories.sql /docker-entrypoint-initdb.d/1004-income-categories.sql
