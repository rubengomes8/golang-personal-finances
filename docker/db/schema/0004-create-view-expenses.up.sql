create view expenses_view as (
	select 
		e.id, e.value, e.date, e.description, es.category_id, ec.name as category_name, 
        e.subcategory_id, es.name as subcategory_name, e.card_id, c.name as card_name 
	from expenses e 
	join cards c on e.card_id = c.id
	join expense_subcategories es on e.subcategory_id = es.id
	join expense_categories ec on ec.id = es.category_id
);