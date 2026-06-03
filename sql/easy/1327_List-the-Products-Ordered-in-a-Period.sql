select product_name, sum(unit) unit
from Orders o
join Products p
on o.product_id = p.product_id
where order_date >= '2020-02-01'::date and order_date < '2020-03-01'::date
group by p.product_id, product_name
having sum(unit) >= 100
