create view incomes_view as (
	select 
		i.id, i.value, i.date, i.description, i.category_id, 
        ic.name as category_name, i.card_id, c.name as card_name 
	from incomes i 
	join cards c on i.card_id = c.id
	join income_categories ic on i.category_id = ic.id
);