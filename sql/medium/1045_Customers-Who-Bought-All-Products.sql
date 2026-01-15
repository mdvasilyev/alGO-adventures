select customer_id
from (select distinct on (customer_id, product_key) * from Customer)
group by customer_id
having count(product_key) = (select count(*) from Product)
